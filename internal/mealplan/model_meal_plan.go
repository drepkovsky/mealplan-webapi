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

type MealPlan struct {

	// Unique meal plan identifier
	Id string `json:"id"`

	// Date for which the meal plan is scheduled
	Date string `json:"date"`

	// Reference to the patient
	PatientId string `json:"patientId"`

	// Meals included in the plan with scheduled times
	Meals []MealPlanMealsInner `json:"meals"`
}