package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.DELETE("/delete-event/:eventId", deleteEvent)
	server.GET("/event/:eventId", getEvent)
	server.PUT("/event/:eventId", updateEvent)
}
