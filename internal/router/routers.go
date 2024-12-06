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

	// remove cors
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Service routes
	serviceGroup := r.Group("/services")
	{
		serviceGroup.POST("", h.CreateService)
		serviceGroup.PUT("/:id", h.UpdateService)
		serviceGroup.DELETE("/:id", h.DeleteService)
		serviceGroup.GET("/:id", h.GetServiceByID)
		serviceGroup.GET("", h.ListServices)

		// Assign group to service
		serviceGroup.POST("/:id/group", h.AssignGroupToService)
	}

	groupGroup := r.Group("/groups")
	{
		groupGroup.GET("/:id", h.GetGroupByID)
		groupGroup.GET("", h.ListGroups)
	}

	// Classification routes
	r.POST("/classify", h.ClassifyService)

	return r
}
