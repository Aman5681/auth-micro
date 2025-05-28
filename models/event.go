package models

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Event struct {
	EventId     uuid.UUID `json:"eventId"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId" binding:"required"`
}

var events = []Event{}

func (e *Event) Save(context *gin.Context) {
	e.EventId = uuid.New()
	e.DateTime = time.Now()
	events = append(events, *e)
	context.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetALLEvents() []Event {
	return events
}
