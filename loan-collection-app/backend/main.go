package main

import (
	"backend/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// For simplicity, use in-memory storage for agent & cases
	agentRoutes := r.Group("/api")
	{
		// Authentication
		agentRoutes.POST("/login", handlers.LoginHandler)

		// Protected routes
		agentRoutes.GET("/cases", handlers.AuthMiddleware, handlers.GetAgentCases)
		agentRoutes.GET("/cases/:caseID", handlers.AuthMiddleware, handlers.GetCaseDetails)
		agentRoutes.POST("/cases/:caseID/trails", handlers.AuthMiddleware, handlers.AddTrail)
		agentRoutes.GET("/cases/:caseID/trails", handlers.AuthMiddleware, handlers.GetTrails)
		agentRoutes.GET("/cases/:caseID/payment-link", handlers.AuthMiddleware, handlers.GetPaymentLink)
	}

	// Serve the React build (if you build the frontend and serve statically)
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	// Optionally serve static files for the React build
	r.StaticFS("/static", http.Dir("./frontend/build/static"))

	log.Println("Starting server on :8080")
	r.Run(":8080")
}
