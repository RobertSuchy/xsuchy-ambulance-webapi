package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/RobertSuchy/xsuchy-ambulance-webapi/api"
    "github.com/RobertSuchy/xsuchy-ambulance-webapi/internal/xsuchy_ambulance"
    "github.com/RobertSuchy/xsuchy-ambulance-webapi/internal/db_service"
    "context"
    "time"
    "github.com/gin-contrib/cors"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("AMBULANCE_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
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
    dbService := db_service.NewMongoService[xsuchy_ambulance.Department](db_service.MongoServiceConfig{Collection: "departments"})
    dbServiceTransport := db_service.NewMongoService[xsuchy_ambulance.Transport](db_service.MongoServiceConfig{Collection: "transports"})
    dbServiceMobilityStatus := db_service.NewMongoService[xsuchy_ambulance.MobilityStatus](db_service.MongoServiceConfig{Collection: "mobility_statuses"})

    defer dbService.Disconnect(context.Background())
    defer dbServiceTransport.Disconnect(context.Background())
    defer dbServiceMobilityStatus.Disconnect(context.Background())

    engine.Use(func(ctx *gin.Context) {
        ctx.Set("db_service_department", dbService)
        ctx.Set("db_service_transport", dbServiceTransport)
        ctx.Set("db_service_mobility_status", dbService)
        ctx.Next()
    })

    // request routings
    xsuchy_ambulance.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}