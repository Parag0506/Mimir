package main

import (
	"github.com/gin-gonic/gin"
	"github.com/your-org/mimir-ai/services/api-gateway-go/internal/handlers"
	"github.com/your-org/mimir-ai/services/api-gateway-go/internal/middleware"
)

func main() {
	router := gin.Default()

	// Apply Middleware
	router.Use(middleware.AuthMiddleware)

	// Define Routes
	router.GET("/health", handlers.HealthCheck)
	router.POST("/api/route", handlers.RouteRequest)

	// Start Server
	router.Run(":8080")
}
