package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEventById)
	server.POST("/events", AddEvent)
}
