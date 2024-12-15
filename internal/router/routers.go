package router

import (
	"backend/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// remove cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// Service routes
	serviceGroup := r.Group("/services")
	{
		serviceGroup.POST("", h.CreateService)
		serviceGroup.GET("", h.ListServices)
		serviceGroup.GET("/:id", h.GetServiceByID)

		serviceGroup.POST("/:id/approve", h.ApproveService)
		serviceGroup.GET("/:id/proposed_classes", h.ListProposedClasses)
	}

	classGroup := r.Group("/classes")
	{
		classGroup.GET("", h.ListClasses)
		classGroup.GET("/:id", h.GetClassByID)
		classGroup.POST("", h.CreateClass)
		classGroup.PUT("/:id", h.UpdateClass)
		classGroup.DELETE("/:id", h.DeleteClass)
	}

	parameterGroup := r.Group("/parameters")
	{
		parameterGroup.POST("", h.CreateParameter)
		parameterGroup.GET("", h.ListParameters)
		parameterGroup.GET("/:id", h.GetParameterByID)
		parameterGroup.PUT("/:id", h.UpdateParameter)
		parameterGroup.DELETE("/:id", h.DeleteParameter)
	}

	r.GET("/report", h.BuildReport)

	return r
}
