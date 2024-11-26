package main

import (
	"calendar-log-backend/config"
	"calendar-log-backend/middlewares"
	"calendar-log-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	routes.RegisterRoutes(r)

	r.Run(config.SERVER_PORT)
}
