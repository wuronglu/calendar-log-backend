package main

import (
	"calendar-log-backend/config"
	"calendar-log-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(config.SERVER_PORT)
}
