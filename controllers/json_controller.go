package controllers

import (
	"calendar-log-backend/config"
	"calendar-log-backend/utils"
	"net/http"
	"os"

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

func CreateJSON(c *gin.Context) {
	var input struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input format",
		})
		return
	}

	data, err := utils.ReadJSONFile(config.JSON_PATH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 检查键是否已经存在
	if _, exists := data[input.Key]; exists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Key already exists",
		})
		return
	}

	// 添加新键值对
	data[input.Key] = input.Value
	if err := utils.WriteJSONFile(config.JSON_PATH, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Key-value pair created successfully",
	})
}

func UpdateJSON(c *gin.Context) {
	key := c.Param("key")
	var input struct {
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input format",
		})
		return
	}

	data, err := utils.ReadJSONFile(config.JSON_PATH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 检查键是否存在
	if _, exists := data[key]; !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Key not found",
		})
		return
	}

	// 更新键值对
	data[key] = input.Value
	if err := utils.WriteJSONFile(config.JSON_PATH, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Key-value pair updated successfully",
	})
}

func DeleteJSON(c *gin.Context) {
	key := c.Param("key")

	data, err := utils.ReadJSONFile(config.JSON_PATH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 检查键是否存在
	if _, exists := data[key]; !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Key not found",
		})
		return
	}

	delete(data, key)

	if err := utils.WriteJSONFile(config.JSON_PATH, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Key deleted successfully",
	})
}

// 下载 JSON 文件的接口
func DownloadJSON(c *gin.Context) {
	// 读取 JSON 文件数据
	data, err := utils.ReadJSONFile(config.JSON_PATH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "data-*.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer os.Remove(tempFile.Name()) // 确保临时文件在响应后被删除

	// 将 JSON 数据写入临时文件
	fileContent, err := utils.MarshalJSON(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 写入文件
	if _, err := tempFile.Write(fileContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 重置文件指针，准备发送文件
	tempFile.Seek(0, 0)

	// 设置响应头，指定文件名和文件类型
	c.Header("Content-Disposition", "attachment; filename=links.json")
	c.Header("Content-Type", "application/json")
	c.File(tempFile.Name())
}
