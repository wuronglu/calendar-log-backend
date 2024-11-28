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
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file not found, using environment variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "9191"
	}

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	routes.RegisterRoutes(r)

	// 启动 HTTPS 服务器
	err = r.RunTLS(":443", "./fullchain.pem", "./privkey.pem")
	if err != nil {
		log.Fatalln("Failed to start server: ", err.Error())
	}
	// r.Run(":" + port)
}
