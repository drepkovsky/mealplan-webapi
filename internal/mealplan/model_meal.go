/*
 * Hospital Meal Plan Management API
 *
 * API to manage meals, patients, and their meal plans in a hospital.
 *
 * API version: 1.0.0
 * Contact: <your_email>
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mealplan

type Meal struct {

	// Unique identifier of the meal
	Id string `json:"id"`

	// Name of the meal
	Name string `json:"name"`

	Nutrients MealNutrients `json:"nutrients"`

	// Size of a single portion
	PortionSize string `json:"portionSize"`

	// List of potential allergens
	Allergens []string `json:"allergens"`

	// List of ingredients used in the meal
	Ingredients []string `json:"ingredients"`
}