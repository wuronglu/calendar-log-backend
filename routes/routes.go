package routes

import (
	"calendar-log-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/map-links", controllers.ReadJSON)           // 查询所有键值对
		api.POST("/map-links", controllers.CreateJSON)        // 创建新键值对
		api.PUT("/map-links/:key", controllers.UpdateJSON)    // 更新已有键值对
		api.DELETE("/map-links/:key", controllers.DeleteJSON) // 删除键值对
	}

}
