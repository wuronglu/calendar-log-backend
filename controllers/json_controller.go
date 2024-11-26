package controllers

import (
	"calendar-log-backend/config"
	"calendar-log-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadJSON(c *gin.Context) {

	data, err := utils.ReadJSONFile(config.JSON_PATH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
