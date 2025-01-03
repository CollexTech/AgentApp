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
	agentRoutes := router.Group("/api")
	agentRoutesV1 := agentRoutes.Group("/v1")
	{
		// Authentication
		agentRoutesV1.POST("/login", handlers.LoginHandler)
		agentRoutesV1.POST("/users/register",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("create_user"),
			handlers.CreateUserHandler)

		agentRoutesV1.GET("/users",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_users"),
			handlers.ListAllUsers)

		// Protected routes
		agentRoutesV1.GET("/cases",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_cases"),
			handlers.GetAgentCases)

		agentRoutesV1.GET("/cases/:caseID",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_cases"),
			handlers.GetCaseDetails)

		agentRoutesV1.POST("/cases/:caseID/trails",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("add_trail"),
			handlers.AddTrail)

		agentRoutesV1.GET("/cases/:caseID/trails",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_trails"),
			handlers.GetTrails)

		agentRoutesV1.GET("/cases/:caseID/payment-link",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("generate_payment_link"),
			handlers.GetPaymentLink)

		// Role Management Routes (Admin Only)
		agentRoutesV1.POST("/roles",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("create_role"),
			handlers.CreateRole)

		agentRoutesV1.PUT("/roles/:id",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("update_role"),
			handlers.UpdateRole)

		agentRoutesV1.DELETE("/roles/:id",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("delete_role"),
			handlers.DeleteRole)

		agentRoutesV1.POST("/roles/assign",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("assign_role_to_user"),
			handlers.AssignRoleToUser)

		agentRoutesV1.POST("/roles/remove",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("remove_role_from_user"),
			handlers.RemoveRoleFromUser)

		agentRoutesV1.GET("/users/:user_id/roles",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_roles"),
			handlers.GetRolesByUser)

		agentRoutesV1.GET("/roles",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_roles"),
			handlers.ListAllRoles)

		agentRoutesV1.GET("/permissions/me",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_my_permissions"),
			handlers.GetMyPermissions)

		agentRoutesV1.GET("/agencies",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_agencies"),
			handlers.ListAllAgencies)

		agentRoutesV1.POST("/agencies",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("create_agency"),
			handlers.CreateAgency)

		agentRoutesV1.DELETE("/agencies/:id",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("delete_agency"),
			handlers.DeleteAgency)

		agentRoutesV1.POST("/agencies/users",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("assign_agency_user"),
			handlers.AssignUserToAgency)

		agentRoutesV1.GET("/agencies/:agency_id/users",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_agency_users"),
			handlers.GetAgencyUsers)

		agentRoutesV1.GET("/agencies/unassigned-users",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_unassigned_users"),
			handlers.GetUnassignedUsers)

		agentRoutesV1.POST("/cases/upload",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("upload_cases"),
			handlers.UploadCasesHandler)

		agentRoutesV1.GET("/cases/unassigned",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("view_unassigned_cases"),
			handlers.GetUnassignedCasesHandler)

		agentRoutesV1.POST("/cases/assign",
			middlewares.AuthMiddleware,
			middlewares.PermissionMiddleware("assign_cases"),
			handlers.AssignCasesHandler)

	}
	router.GET("/health-check", handlers.Healthcheck)
	router.OPTIONS("/*any", handlers.OptionsHandle)
}
