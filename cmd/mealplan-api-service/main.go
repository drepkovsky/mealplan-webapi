package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/drepkovsky/mealplan-webapi/api"
	"github.com/drepkovsky/mealplan-webapi/internal/db_service"
	"github.com/drepkovsky/mealplan-webapi/internal/mealplan"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("MEALPLAN_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("MEALPLAN_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())
    corsMiddleware := cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{""},
        AllowCredentials: false,
        MaxAge: 12 * time.Hour,
    })
    engine.Use(corsMiddleware)

     // setup context update  middleware
     dbService := db_service.NewMongoService[mealplan.Ambulance](db_service.MongoServiceConfig{})
     defer dbService.Disconnect(context.Background())
     engine.Use(func(ctx *gin.Context) {
         ctx.Set("db_service", dbService)
         ctx.Next()
     })

    // request routings
    mealplan.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}