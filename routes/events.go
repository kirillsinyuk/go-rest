package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/models"
	"strconv"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEventById(context *gin.Context) {
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

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusCreated, event)
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	models.DeleteById(id)
	context.JSON(http.StatusOK, gin.H{})
}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	_, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
