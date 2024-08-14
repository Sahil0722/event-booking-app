package main

import (
	"github.com/Sahil0722/event-booking-app/db"
	"github.com/Sahil0722/event-booking-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
