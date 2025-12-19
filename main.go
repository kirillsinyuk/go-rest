package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/models"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", addEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func addEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.Save()
	context.JSON(http.StatusCreated, event)
}
