package routes

import (
	"calendar-log-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/map-links", controllers.ReadJSON)
		api.POST("/map-links", controllers.CreateJSON)
		api.PUT("/map-links/:key", controllers.UpdateJSON)
		api.DELETE("/map-links/:key", controllers.DeleteJSON)
		api.GET("/map-links/download", controllers.DownloadJSON)
	}

}
