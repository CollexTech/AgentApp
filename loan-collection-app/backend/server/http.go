package server

import (
	"backend/handlers"
	"backend/handlers/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initialize() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(gin.CustomRecovery(middlewares.ErrorHandler))
	route(&router.RouterGroup)
	router.Run(fmt.Sprintf(":%s", viper.GetString("PORT")))
}

func route(router *gin.RouterGroup) {
	zapLogger, _ := zap.NewProductionConfig().Build()
	validator := validator.New()
	decoder := schema.NewDecoder()
	router.Use(middlewares.EnvMiddleware(zapLogger, validator, decoder))
	router.Use(middlewares.SetCorsHeaders())
	agentRoutes := router.Group("/agent")
	agentBackendRoutes := agentRoutes.Group("/api")
	agentRoutesV1 := agentBackendRoutes.Group("/v1")
	{
		// Authentication
		agentRoutesV1.POST("/login", handlers.LoginHandler)
		agentRoutesV1.POST("/register", handlers.CreateUserHandler)

		// Protected routes
		agentRoutesV1.GET("/cases", middlewares.AuthMiddleware, handlers.GetAgentCases)
		agentRoutesV1.GET("/cases/:caseID", middlewares.AuthMiddleware, handlers.GetCaseDetails)
		agentRoutesV1.POST("/cases/:caseID/trails", middlewares.AuthMiddleware, handlers.AddTrail)
		agentRoutesV1.GET("/cases/:caseID/trails", middlewares.AuthMiddleware, handlers.GetTrails)
		agentRoutesV1.GET("/cases/:caseID/payment-link", middlewares.AuthMiddleware, handlers.GetPaymentLink)
	}
	router.GET("/health-check", handlers.Healthcheck)
	router.OPTIONS("/*any", handlers.OptionsHandle)
}
