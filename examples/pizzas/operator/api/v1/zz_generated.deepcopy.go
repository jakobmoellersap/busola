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
	out.Status = in.Status
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
	if in.Pizzas != nil {
		in, out := &in.Pizzas, &out.Pizzas
		*out = make(Pizzas, len(*in))
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
	if in.Toppings != nil {
		in, out := &in.Toppings, &out.Toppings
		*out = make(Toppings, len(*in))
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
func (in Pizzas) DeepCopyInto(out *Pizzas) {
	{
		in := &in
		*out = make(Pizzas, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pizzas.
func (in Pizzas) DeepCopy() Pizzas {
	if in == nil {
		return nil
	}
	out := new(Pizzas)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topping) DeepCopyInto(out *Topping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topping.
func (in *Topping) DeepCopy() *Topping {
	if in == nil {
		return nil
	}
	out := new(Topping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Toppings) DeepCopyInto(out *Toppings) {
	{
		in := &in
		*out = make(Toppings, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Toppings.
func (in Toppings) DeepCopy() Toppings {
	if in == nil {
		return nil
	}
	out := new(Toppings)
	in.DeepCopyInto(out)
	return *out
}
