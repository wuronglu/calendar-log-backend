package main

import (
	"calendar-log-backend/middlewares"
	"calendar-log-backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9191"
	}

	r.Use(middlewares.CORSMiddleware())
	routes.RegisterRoutes(r)

	r.Run(":" + port)
}
