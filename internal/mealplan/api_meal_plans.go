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

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MealPlansAPI interface {

	// internal registration of api routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateMealPlan - Creates a new meal plan
	CreateMealPlan(ctx *gin.Context)

	// DeleteMealPlan - Delete a meal plan by ID
	DeleteMealPlan(ctx *gin.Context)

	// GetMealPlan - Get a meal plan by ID
	GetMealPlan(ctx *gin.Context)

	// ListMealPlans - List all meal plans for given patient
	ListMealPlans(ctx *gin.Context)

	// UpdateMealPlan - Update a meal plan by ID
	UpdateMealPlan(ctx *gin.Context)
}

// partial implementation of MealPlansAPI - all functions must be implemented in add on files
type implMealPlansAPI struct {
}

func newMealPlansAPI() MealPlansAPI {
	return &implMealPlansAPI{}
}

func (this *implMealPlansAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/meal-plans", this.CreateMealPlan)
	routerGroup.Handle(http.MethodDelete, "/meal-plans/:mealPlanId", this.DeleteMealPlan)
	routerGroup.Handle(http.MethodGet, "/meal-plans/:mealPlanId", this.GetMealPlan)
	routerGroup.Handle(http.MethodGet, "/meal-plans/patient/:patientId", this.ListMealPlans)
	routerGroup.Handle(http.MethodPut, "/meal-plans/:mealPlanId", this.UpdateMealPlan)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // CreateMealPlan - Creates a new meal plan
// func (this *implMealPlansAPI) CreateMealPlan(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteMealPlan - Delete a meal plan by ID
// func (this *implMealPlansAPI) DeleteMealPlan(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetMealPlan - Get a meal plan by ID
// func (this *implMealPlansAPI) GetMealPlan(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // ListMealPlans - List all meal plans for given patient
// func (this *implMealPlansAPI) ListMealPlans(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateMealPlan - Update a meal plan by ID
// func (this *implMealPlansAPI) UpdateMealPlan(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
