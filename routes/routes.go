package routes

import (
	"github.com/gin-gonic/gin"
	"rest/auth"
)

func SetupRoutes(server *gin.Engine) {

	authGroup := server.Group("/")
	authGroup.Use(auth.Authentificate)
	authGroup.POST("/events", createEvent)
	authGroup.PUT("/events/:id", updateEvent)
	authGroup.DELETE("/events/:id", deleteEvent)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
