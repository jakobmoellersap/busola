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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PizzaOrderSpec defines the desired state of PizzaOrder
type PizzaOrderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Order description
	Description string `json:"description,omitempty"`

	// The details of the order
	OrderDetails string `json:"orderDetails,omitempty"`

	// Ordered pizzas
	Pizzas `json:"pizzas,omitempty"`
}

// Ordered pizzas
type Pizzas []Pizza

// PizzaOrderStatus defines the observed state of PizzaOrder
type PizzaOrderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State `json:"state,omitempty"`
}

// The details of the order
type OrderDetails struct {
	// The details of the order
	PaymentMethod `json:"paymentMethod,omitempty"`
	// The details of the order
	Realization `json:"realization,omitempty"`
}

type PaymentMethod string

const CardPaymentMethod PaymentMethod = "CARD"
const CashPaymentMethod PaymentMethod = "CASH"

type Realization string

const SelfPickupRealization Realization = "SELF-PICKUP"
const DeliveryRealization Realization = "DELIVERY"

// +kubebuilder:validation:Enum=Processing;Deleting;Ready;Error
type State string

// Valid Kyma States.
const (
	Ready      State = "Ready"
	Processing State = "Processing"
	Error      State = "Error"
	Deleting   State = "Deleting"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PizzaOrder is the Schema for the pizzaorders API
type PizzaOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PizzaOrderSpec   `json:"spec,omitempty"`
	Status PizzaOrderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PizzaOrderList contains a list of PizzaOrder
type PizzaOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PizzaOrder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PizzaOrder{}, &PizzaOrderList{})
}
