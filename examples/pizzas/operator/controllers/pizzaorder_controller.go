/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	busolaexamplecomv1 "github.com/kyma-project/busola/examples/pizzas/operator/api/v1"
	v1 "github.com/kyma-project/busola/examples/pizzas/operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ErrStillCooking = errors.New("still cooking some pizzas!")

// PizzaOrderReconciler reconciles a PizzaOrder object
type PizzaOrderReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	record.EventRecorder
}

//+kubebuilder:rbac:groups=busola.example.com,resources=pizzaorders,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=busola.example.com,resources=pizzaorders/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=busola.example.com,resources=pizzaorders/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=events,verbs=create;patch;get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PizzaOrder object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *PizzaOrderReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// check if kyma resource exists
	order := &v1.PizzaOrder{}
	if err := r.Get(ctx, req.NamespacedName, order); err != nil {
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		logger.Info(req.NamespacedName.String() + " got deleted!")
		return ctrl.Result{}, client.IgnoreNotFound(err) //nolint:wrapcheck
	}

	// check if deletionTimestamp is set, retry until it gets fully deleted
	if !order.DeletionTimestamp.IsZero() && order.Status.State != v1.Deleting {
		order.Status.State = v1.Deleting
		// if the status is not yet set to deleting, also update the status
		if err := r.Status().Update(ctx, order); err != nil {
			return ctrl.Result{RequeueAfter: 3 * time.Second}, fmt.Errorf(
				"could not update status after triggering deletion: %w", err)
		}
		return ctrl.Result{}, nil
	}

	if controllerutil.ContainsFinalizer(order, v1.PizzaOrderFinalizer) {
		controllerutil.AddFinalizer(order, v1.PizzaOrderFinalizer)
		if err := r.Update(ctx, order); err != nil {
			return ctrl.Result{RequeueAfter: 3 * time.Second},
				fmt.Errorf("could not update order after finalizer check: %w", err)
		}
		return ctrl.Result{}, nil
	}

	// state handling
	return r.stateHandling(ctx, order)
}

// SetupWithManager sets up the controller with the Manager.
func (r *PizzaOrderReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&busolaexamplecomv1.PizzaOrder{}).
		Owns(&busolaexamplecomv1.Pizza{}).
		Complete(r)
}

func (r *PizzaOrderReconciler) stateHandling(ctx context.Context, order *v1.PizzaOrder) (ctrl.Result, error) {
	switch order.Status.State {
	case "":
		return ctrl.Result{}, r.HandleInitialState(ctx, order)
	case v1.Processing:
		processResult := r.HandleProcessingState(ctx, order)
		if processResult == ErrStillCooking {
			return ctrl.Result{RequeueAfter: 1 * time.Second}, nil
		}
		return ctrl.Result{}, nil
	case v1.Deleting:
		return ctrl.Result{}, r.HandleDeletingState(ctx, order)
	case v1.Error:
		return ctrl.Result{}, r.HandleErrorState(ctx, order)
	case v1.Ready:
		return ctrl.Result{}, r.HandleReadyState(ctx, order)
	}

	return ctrl.Result{}, nil
}

func (r *PizzaOrderReconciler) HandleInitialState(ctx context.Context, order *v1.PizzaOrder) error {
	order.Status.State = v1.Processing
	order.Status.ObservedGeneration = order.Generation
	// if the status is not yet set to deleting, also update the status
	if err := r.Status().Update(ctx, order); err != nil {
		return fmt.Errorf("could not update kyma status after triggering processing: %w", err)
	}
	r.EventRecorder.Event(order, "Normal", "NewOrder", "order was picked up by the pizza-operator")
	return nil
}

func (r *PizzaOrderReconciler) HandleProcessingState(ctx context.Context, order *v1.PizzaOrder) error {
	withStateUpdateOnErr := func(err error) error {
		order.Status.State = v1.Error
		order.Status.ObservedGeneration = order.Generation
		if statusErr := r.Status().Update(ctx, order); err != nil {
			return fmt.Errorf("could not update kyma status: %w (%s)", err, statusErr.Error())
		}
		return nil
	}

	templatedPizzas := make(map[*v1.PizzaTemplate]*v1.Pizza)
	for _, pizzaTemplate := range order.Spec.PizzaTemplates {
		pizzas := &v1.PizzaList{}
		// pizzaTemplate.Selector
		ls, err := metav1.LabelSelectorAsSelector(&pizzaTemplate.Selector)

		if err != nil {
			return withStateUpdateOnErr(err)
		}

		err = r.List(ctx, pizzas, client.MatchingLabelsSelector{Selector: ls})

		if err != nil {
			return withStateUpdateOnErr(err)
		}

		templatedPizzas[&pizzaTemplate] = &pizzas.Items[0]
	}

	order.Status.Bill = v1.Bill{
		Items: []v1.Billable{},
		Sum:   "",
	}

	var sum USD = 0
	for template, pizza := range templatedPizzas {
		var price USD
		for _, pricingEntry := range pizza.Spec.PricingTable {
			if pricingEntry.Size == template.Size {
				price, _ = NewUSD(pricingEntry.Price)
				price = price.Multiply(float64(template.Quantity))
				break
			}
		}
		for _, ingredient := range pizza.Spec.Ingredients {
			ingredientPrice, _ := NewUSD(ingredient.Price)
			if ingredient.Quantity == 0 {
				ingredient.Quantity = 1
			}
			ingredientPrice = ingredientPrice.Multiply(float64(ingredient.Quantity))
			price += ingredientPrice
		}
		for _, selectedTopping := range template.SelectedToppings {
			for _, topping := range pizza.Spec.Toppings {
				if selectedTopping == topping.Name {
					toppingPrice, _ := NewUSD(topping.Price)
					if topping.Quantity == 0 {
						topping.Quantity = 1
					}
					toppingPrice = toppingPrice.Multiply(float64(topping.Quantity))
					price += toppingPrice
				}
			}
		}
		name := fmt.Sprintf("%s with %s", pizza.Name, strings.Join(template.SelectedToppings, ","))
		order.Status.Bill.Items = append(order.Status.Bill.Items, v1.Billable{
			Name:     name,
			Quantity: template.Quantity,
			Price:    price.String(),
		})
		sum += price

		alreadyInOven := false
		for _, pizzaInOven := range order.Status.Oven.Pizzas {
			if pizzaInOven.Name == name {
				alreadyInOven = true
				break
			}
		}
		if !alreadyInOven {
			pizzaInOven := v1.PizzaInOven{
				Name:          name,
				PutIntoOvenAt: time.Now().Format(time.RFC3339),
				Done:          false,
				CookingTime:   pizza.Spec.CookingTime,
			}
			r.EventRecorder.Eventf(order, "Normal", "NewPizza", "put a fresh pizza into the oven: %s (for %s)", pizzaInOven.Name, pizzaInOven.CookingTime)
			order.Status.Oven.Pizzas = append(order.Status.Oven.Pizzas, pizzaInOven)
		}
	}
	order.Status.Bill.Sum = sum.String()

	doneCookingPizzas := true
	for _, pizzaInOven := range order.Status.Oven.Pizzas {
		cookingTime, _ := time.ParseDuration(pizzaInOven.CookingTime)
		putIntoOvenAt, _ := time.Parse(time.RFC3339, pizzaInOven.PutIntoOvenAt)
		// still cooking
		if time.Since(putIntoOvenAt) < cookingTime {
			doneCookingPizzas = false
			continue
		}
		pizzaInOven.Done = true
		r.EventRecorder.Eventf(order, "Normal", "PizzaDone", "the pizza is finished: %s", pizzaInOven)
	}
	if !doneCookingPizzas {
		if err := r.Status().Update(ctx, order); err != nil {
			return fmt.Errorf("could not update kyma status after triggering processing: %w", err)
		}
		return ErrStillCooking
	}

	order.Status.State = v1.Ready
	order.Status.ObservedGeneration = order.Generation
	// if the status is not yet set to deleting, also update the status
	if err := r.Status().Update(ctx, order); err != nil {
		return fmt.Errorf("could not update kyma status after triggering processing: %w", err)
	}
	r.EventRecorder.Eventf(order, "Normal", "OrderReady", "the order is finished!")
	return nil
}

func (r *PizzaOrderReconciler) HandleDeletingState(ctx context.Context, order *v1.PizzaOrder) error {
	controllerutil.RemoveFinalizer(order, v1.PizzaOrderFinalizer)
	if err := r.Update(ctx, order); err != nil {
		return fmt.Errorf("error while trying to udpate kyma during deletion: %w", err)
	}
	return nil
}

func (r *PizzaOrderReconciler) HandleErrorState(ctx context.Context, order *v1.PizzaOrder) error {
	return r.HandleConsistencyChanges(ctx, order)
}

func (r *PizzaOrderReconciler) HandleReadyState(ctx context.Context, order *v1.PizzaOrder) error {
	return r.HandleConsistencyChanges(ctx, order)
}

func (r *PizzaOrderReconciler) HandleConsistencyChanges(ctx context.Context, order *v1.PizzaOrder) error {
	if order.Status.ObservedGeneration != order.Generation {
		order.Status.State = v1.Processing
		order.Status.ObservedGeneration = order.Generation
		// if the status is not yet set to deleting, also update the status
		if err := r.Status().Update(ctx, order); err != nil {
			return fmt.Errorf("could not update kyma status after triggering processing: %w", err)
		}
	}
	return nil
}
