package mealplan

import (
	"net/http"

	"github.com/drepkovsky/mealplan-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateMealPlan - Creates a new meal plan
func (this *implMealPlansAPI) CreateMealPlan(ctx *gin.Context) {
	value, exists := ctx.Get("meal_plan_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[MealPlan])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	mealPlan := MealPlan{}
	err := ctx.BindJSON(&mealPlan)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "invalid meal plan data",
				"error":   err.Error(),
			})
		return
	}

	if mealPlan.Id != "" {
		mealPlan.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, mealPlan.Id, &mealPlan)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			mealPlan,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Meal Plan already exists in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create meal plan in database",
				"error":   err.Error(),
			},
		)
	}
}

// DeleteMealPlan - Delete a meal plan by ID
func (this *implMealPlansAPI) DeleteMealPlan(ctx *gin.Context) {
	value, exists := ctx.Get("meal_plan_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[MealPlan])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	mealPlanID := ctx.Param("mealPlanId")
	err := db.DeleteDocument(ctx, mealPlanID)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Meal Plan not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete meal plan in database",
				"error":   err.Error(),
			},
		)
	}
}

// GetMealPlan - Get a meal plan by ID
func (this *implMealPlansAPI) GetMealPlan(ctx *gin.Context) {
	value, exists := ctx.Get("meal_plan_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[MealPlan])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	mealPlanID := ctx.Param("mealPlanId")
	mealPlan, err := db.FindDocument(ctx, bson.D{{Key: "id", Value: mealPlanID}})

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			mealPlan,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Meal Plan not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to get meal plan from database",
				"error":   err.Error(),
			},
		)
	}
}

// ListMealPlans - List all meal plans
func (this *implMealPlansAPI) ListMealPlans(ctx *gin.Context) {
	value, exists := ctx.Get("meal_plan_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[MealPlan])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	patientID := ctx.Param("patientId")
	mealPlans, err := db.ListDocuments(ctx, bson.D{{Key: "patientid", Value: patientID}})

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			mealPlans,
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to list meal plans from database",
				"error":   err.Error(),
			},
		)
	}
}

// UpdateMealPlan - Update a meal plan by ID
func (this *implMealPlansAPI) UpdateMealPlan(ctx *gin.Context) {
	value, exists := ctx.Get("meal_plan_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[MealPlan])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	mealPlanID := ctx.Param("mealPlanId")
	mealPlan := MealPlan{}
	err := ctx.BindJSON(&mealPlan)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "invalid meal plan data",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, mealPlanID, &mealPlan)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			mealPlan,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Meal Plan not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update meal plan in database",
				"error":   err.Error(),
			},
		)
	}
}
