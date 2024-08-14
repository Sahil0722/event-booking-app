package routes

import (
	"github.com/Sahil0722/event-booking-app/middlewares"
	"github.com/Sahil0722/event-booking-app/services/events"
	"github.com/Sahil0722/event-booking-app/services/registers"
	"github.com/Sahil0722/event-booking-app/services/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events Endpoints
	server.GET("/events", events.GetEvents)
	server.GET("/events/:id", events.GetEvent)
	server.GET("/users", users.GetUsers)
	server.GET("/registers", registers.GetRegistrations)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authorize)
	authenticated.POST("/events", events.CreateEvent)
	authenticated.PUT("/events/:id", events.UpdateEvent)
	authenticated.DELETE("/events/:id", events.DeleteEvent)

	authenticated.POST("/events/:id/register", registers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", registers.CancelRegistrationEvent)

	// User Endpoints
	server.POST("/signup", users.SignUp)
	server.POST("login", users.Login)
}
