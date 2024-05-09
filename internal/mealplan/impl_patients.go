package mealplan

import (
	"net/http"

	"github.com/drepkovsky/mealplan-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// CreatePatient - Creates a new patient
func (this *implPatientsAPI) CreatePatient(ctx *gin.Context) {
	value, exists := ctx.Get("patient_service")
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

	db, ok := value.(db_service.DbService[Patient])
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

	patient := Patient{}
	err := ctx.BindJSON(&patient)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "invalid patient data",
				"error":   err.Error(),
			})
		return
	}

	if patient.PatientId != "" {
		patient.PatientId = uuid.New().String()
	}

	err = db.CreateDocument(ctx, patient.PatientId, &patient)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			patient,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Patient already exists in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create patient in database",
				"error":   err.Error(),
			},
		)
	}
}

// DeletePatient - Delete a patient by ID
func (this *implPatientsAPI) DeletePatient(ctx *gin.Context) {
	value, exists := ctx.Get("patient_service")
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

	db, ok := value.(db_service.DbService[Patient])
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
	err := db.DeleteDocument(ctx, patientID)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete patient in database",
				"error":   err.Error(),
			},
		)
	}
}

// GetPatient - Get a patient by ID
func (this *implPatientsAPI) GetPatient(ctx *gin.Context) {
	value, exists := ctx.Get("patient_service")
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

	db, ok := value.(db_service.DbService[Patient])
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
	patient, err := db.FindDocument(ctx, bson.D{{Key: "patientid", Value: patientID}})

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			patient,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to get patient from database",
				"error":   err.Error(),
			},
		)
	}
}

// ListPatients - List all patients
func (this *implPatientsAPI) ListPatients(ctx *gin.Context) {
	value, exists := ctx.Get("patient_service")
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

	db, ok := value.(db_service.DbService[Patient])
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

	patients, err := db.ListDocuments(ctx, bson.D{})

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			patients,
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to list patients from database",
				"error":   err.Error(),
			},
		)
	}
}

// UpdatePatient - Update a patient by ID
func (this *implPatientsAPI) UpdatePatient(ctx *gin.Context) {
	value, exists := ctx.Get("patient_service")
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

	db, ok := value.(db_service.DbService[Patient])
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
	patient := Patient{}
	err := ctx.BindJSON(&patient)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "invalid patient data",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, patientID, &patient)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			patient,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Patient not found in database",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update patient in database",
				"error":   err.Error(),
			},
		)
	}
}
