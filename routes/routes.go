package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	events := server.Group("/events")
	events.GET("/", getEvents)
	events.POST("/", createEvent)
	events.DELETE("/delete-event-by-id/:eventId", deleteEvent)
	events.GET("/get-event-by-id/:eventId", getEvent)
	events.PUT("/update-event-by-id/:eventId", updateEvent)

	users := server.Group("/users")
	users.POST("/create-user", signUp)
	users.POST("/login", login)
}
