//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bill) DeepCopyInto(out *Bill) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Billable, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bill.
func (in *Bill) DeepCopy() *Bill {
	if in == nil {
		return nil
	}
	out := new(Bill)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Billable) DeepCopyInto(out *Billable) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Billable.
func (in *Billable) DeepCopy() *Billable {
	if in == nil {
		return nil
	}
	out := new(Billable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Comments) DeepCopyInto(out *Comments) {
	{
		in := &in
		*out = make(Comments, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Comments.
func (in Comments) DeepCopy() Comments {
	if in == nil {
		return nil
	}
	out := new(Comments)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingredient) DeepCopyInto(out *Ingredient) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingredient.
func (in *Ingredient) DeepCopy() *Ingredient {
	if in == nil {
		return nil
	}
	out := new(Ingredient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderDetails) DeepCopyInto(out *OrderDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderDetails.
func (in *OrderDetails) DeepCopy() *OrderDetails {
	if in == nil {
		return nil
	}
	out := new(OrderDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Oven) DeepCopyInto(out *Oven) {
	*out = *in
	if in.Pizzas != nil {
		in, out := &in.Pizzas, &out.Pizzas
		*out = make([]PizzaInOven, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Oven.
func (in *Oven) DeepCopy() *Oven {
	if in == nil {
		return nil
	}
	out := new(Oven)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pizza) DeepCopyInto(out *Pizza) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pizza.
func (in *Pizza) DeepCopy() *Pizza {
	if in == nil {
		return nil
	}
	out := new(Pizza)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pizza) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaInOven) DeepCopyInto(out *PizzaInOven) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaInOven.
func (in *PizzaInOven) DeepCopy() *PizzaInOven {
	if in == nil {
		return nil
	}
	out := new(PizzaInOven)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaList) DeepCopyInto(out *PizzaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pizza, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaList.
func (in *PizzaList) DeepCopy() *PizzaList {
	if in == nil {
		return nil
	}
	out := new(PizzaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrder) DeepCopyInto(out *PizzaOrder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrder.
func (in *PizzaOrder) DeepCopy() *PizzaOrder {
	if in == nil {
		return nil
	}
	out := new(PizzaOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaOrder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderList) DeepCopyInto(out *PizzaOrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PizzaOrder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderList.
func (in *PizzaOrderList) DeepCopy() *PizzaOrderList {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PizzaOrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderSpec) DeepCopyInto(out *PizzaOrderSpec) {
	*out = *in
	out.OrderDetails = in.OrderDetails
	if in.PizzaTemplates != nil {
		in, out := &in.PizzaTemplates, &out.PizzaTemplates
		*out = make(PizzaTemplates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderSpec.
func (in *PizzaOrderSpec) DeepCopy() *PizzaOrderSpec {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaOrderStatus) DeepCopyInto(out *PizzaOrderStatus) {
	*out = *in
	in.Bill.DeepCopyInto(&out.Bill)
	in.Oven.DeepCopyInto(&out.Oven)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaOrderStatus.
func (in *PizzaOrderStatus) DeepCopy() *PizzaOrderStatus {
	if in == nil {
		return nil
	}
	out := new(PizzaOrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaSpec) DeepCopyInto(out *PizzaSpec) {
	*out = *in
	if in.PricingTable != nil {
		in, out := &in.PricingTable, &out.PricingTable
		*out = make(PricingTable, len(*in))
		copy(*out, *in)
	}
	if in.Ingredients != nil {
		in, out := &in.Ingredients, &out.Ingredients
		*out = make([]Ingredient, len(*in))
		copy(*out, *in)
	}
	if in.Toppings != nil {
		in, out := &in.Toppings, &out.Toppings
		*out = make([]Ingredient, len(*in))
		copy(*out, *in)
	}
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = make(Comments, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaSpec.
func (in *PizzaSpec) DeepCopy() *PizzaSpec {
	if in == nil {
		return nil
	}
	out := new(PizzaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaStatus) DeepCopyInto(out *PizzaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaStatus.
func (in *PizzaStatus) DeepCopy() *PizzaStatus {
	if in == nil {
		return nil
	}
	out := new(PizzaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PizzaTemplate) DeepCopyInto(out *PizzaTemplate) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.SelectedToppings != nil {
		in, out := &in.SelectedToppings, &out.SelectedToppings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaTemplate.
func (in *PizzaTemplate) DeepCopy() *PizzaTemplate {
	if in == nil {
		return nil
	}
	out := new(PizzaTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PizzaTemplates) DeepCopyInto(out *PizzaTemplates) {
	{
		in := &in
		*out = make(PizzaTemplates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PizzaTemplates.
func (in PizzaTemplates) DeepCopy() PizzaTemplates {
	if in == nil {
		return nil
	}
	out := new(PizzaTemplates)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pricing) DeepCopyInto(out *Pricing) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pricing.
func (in *Pricing) DeepCopy() *Pricing {
	if in == nil {
		return nil
	}
	out := new(Pricing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PricingTable) DeepCopyInto(out *PricingTable) {
	{
		in := &in
		*out = make(PricingTable, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PricingTable.
func (in PricingTable) DeepCopy() PricingTable {
	if in == nil {
		return nil
	}
	out := new(PricingTable)
	in.DeepCopyInto(out)
	return *out
}