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

// Information about our pizza
type PizzaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Pizza description
	Description string `json:"description,omitempty"`

	// The name of the sauce to use on our pizza
	Sauce `json:"sauce,omitempty"`

	// A Pricing Table for the Pizza
	PricingTable `json:"pricing,omitempty"`

	// List of ingredients for our pizza
	Ingredients []Ingredient `json:"ingredients,omitempty"`

	// List of toppings for our pizza
	Toppings []Ingredient `json:"toppings,omitempty"`

	CookingTime string `json:"cookingTime,omitempty"`

	// Additional comments
	Comments `json:"comments,omitempty"`

	// Additional comments
	RecipeSecret string `json:"recipeSecret,omitempty"`
}

type PricingTable []Pricing
type Pricing struct {
	Size  string `json:"size,omitempty"`
	Price string `json:"price,omitempty"`
}

type Sauce string

const GarlicSauce Sauce = "GARLIC"
const TomatoSauce Sauce = "TOMATO"

type Ingredient struct {
	// The name of the topping
	Name     string `json:"name,omitempty"`
	Price    string `json:"price,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

type Comments []string

// PizzaStatus defines the observed state of Pizza
type PizzaStatus struct {
	// State is the state of the Pizza
	State `json:"state,omitempty"`

	// Observed generation
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pizza is the Schema for the pizzas API
type Pizza struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PizzaSpec   `json:"spec,omitempty"`
	Status PizzaStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PizzaList contains a list of Pizza
type PizzaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pizza `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pizza{}, &PizzaList{})
}