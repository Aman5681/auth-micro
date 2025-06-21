package routes

import (
	"net/http"

	"github.com/Aman5681/auth-micro/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getEvents(context *gin.Context) {
	events, err := models.GetALLEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
	}
	context.JSON(http.StatusOK, events)
}

func updateEvent(context *gin.Context) {
	eventId := context.Param("eventId")
	var updatedEvent models.Event
	var err error
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}
	if eventId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event ID is required"})
	}

	updatedEvent.EventId = eventId

	event, err := models.UpdateEvent(&updatedEvent, context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to update event"})
	}

	context.JSON(http.StatusOK, event)
}

func getEvent(context *gin.Context) {
	eventId := context.Param("eventId")

	if eventId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "eventId is required"})
		return
	}

	if _, err := uuid.Parse(eventId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid eventId format (expected UUID)"})
		return
	}

	event, err := models.GetEventById(&eventId, context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind json"})
		return
	}
	err = event.SaveEvent(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save event"})
	}
}

func deleteEvent(context *gin.Context) {
	eventId := context.Param("eventId")
	if eventId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "eventId is required"})
		return
	}

	if _, err := uuid.Parse(eventId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid eventId format (expected UUID)"})
		return
	}

	err := models.DeleteEvent(&eventId, context)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
