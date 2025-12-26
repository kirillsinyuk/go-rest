package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/models"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event.UserId = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusCreated, event)
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = models.DeleteById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	userId := context.GetInt64("userId")

	if userId != event.UserId {
		context.JSON(http.StatusForbidden, gin.H{"error": "the user does not belong to this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	updatedEvent.ID = id

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, updatedEvent)
}
