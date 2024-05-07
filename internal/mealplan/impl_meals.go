package mealplan

import (
	"net/http"

	"github.com/drepkovsky/mealplan-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateMeal - Creates a new meal
func (this *implMealsAPI) CreateMeal(ctx *gin.Context) {
	value, exists := ctx.Get("meal_service")
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

	db, ok := value.(db_service.DbService[Meal])
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

	meal := Meal{}
	err := ctx.BindJSON(&meal)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "invalid meal data",
				"error":   err.Error(),
			})
		return
	}

	if meal.Id != "" {
		meal.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, meal.Id, &meal)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			meal,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Meal already exists in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create meal in database",
				"error":   err.Error(),
			},
		)
	}

}

// DeleteMeal - Delete a meal by ID
func (this *implMealsAPI) DeleteMeal(ctx *gin.Context) {
	value, exists := ctx.Get("meal_service")
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

	db, ok := value.(db_service.DbService[Meal])
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

	mealID := ctx.Param("mealId")
	err := db.DeleteDocument(ctx, mealID)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Meal not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete meal in database",
				"error":   err.Error(),
			},
		)
	}
}

// GetMeal - Get a meal by ID
func (this *implMealsAPI) GetMeal(ctx *gin.Context) {
	value, exists := ctx.Get("meal_service")
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

	db, ok := value.(db_service.DbService[Meal])
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

	mealID := ctx.Param("mealId")
	meal, err := db.FindDocument(ctx, mealID)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			meal,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Meal not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to get meal from database",
				"error":   err.Error(),
			},
		)
	}
}

// ListMeals - List all meals
func (this *implMealsAPI) ListMeals(ctx *gin.Context) {
	value, exists := ctx.Get("meal_service")
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

	db, ok := value.(db_service.DbService[Meal])
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

	meals, err := db.ListDocuments(ctx, bson.D{})

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			meals,
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to list meals from database",
				"error":   err.Error(),
			},
		)
	}
}

// UpdateMeal - Update a meal by ID
func (this *implMealsAPI) UpdateMeal(ctx *gin.Context) {
	value, exists := ctx.Get("meal_service")
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

	db, ok := value.(db_service.DbService[Meal])
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

	mealID := ctx.Param("mealId")
	meal := Meal{}
	err := ctx.BindJSON(&meal)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "invalid meal data",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, mealID, &meal)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			meal,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Meal not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update meal in database",
				"error":   err.Error(),
			},
		)
	}
}
