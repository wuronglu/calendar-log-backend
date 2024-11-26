package main

import (
	"calendar-log-backend/middlewares"
	"calendar-log-backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9191"
	}

	r.Use(middlewares.CORSMiddleware())
	routes.RegisterRoutes(r)

	r.Run(":" + port)
}
