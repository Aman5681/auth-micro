package main

import (
	"net/http"

	"github.com/Aman5681/auth-micro/db"
	"github.com/Aman5681/auth-micro/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func getEvents(context *gin.Context) {
	events := models.GetALLEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.Save(context)
}
