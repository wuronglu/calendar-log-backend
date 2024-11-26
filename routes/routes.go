package routes

import (
	"calendar-log-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/map-links", controllers.ReadJSON)
	}
}
