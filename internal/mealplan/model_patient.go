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

type Patient struct {

	// Full name of the patient
	FullName string `json:"fullName"`

	// Unique identifier for the patient
	Id string `json:"id"`

	// Age of the patient
	Age int32 `json:"age"`

	// List of patient's allergens
	Allergens []string `json:"allergens"`

	// List of preferred ingredients for the patient
	IngredientPreferences []string `json:"ingredientPreferences"`
}
