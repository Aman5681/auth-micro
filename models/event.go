package models

import (
	"net/http"
	"time"

	"github.com/Aman5681/auth-micro/db"
	"github.com/Aman5681/auth-micro/utils"
	"github.com/gin-gonic/gin"
)

type Event struct {
	EventId     string    `json:"eventId"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserId      string    `json:"userId" binding:"required"`
}

func (e *Event) SaveEvent(context *gin.Context) error {
	e.EventId = utils.GenerateUUID()
	e.DateTime = time.Now()
	query := `
	INSERT INTO events(eventId, name, description, location, dateTime, userId)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.EventId, e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, gin.H{"status": "success"})
	return nil
}

func GetALLEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var events []Event

	for row.Next() {
		var event Event
		err := row.Scan(&event.EventId, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(eventId *string, context *gin.Context) (*Event, error) {
	query := "SELECT * FROM events WHERE eventId = ?"
	row := db.DB.QueryRow(query, &eventId)
	var event Event
	err := row.Scan(&event.EventId, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func DeleteEvent(eventId *string, context *gin.Context) error {
	query := "DELETE FROM events WHERE eventId = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&eventId)
	if err != nil {
		return err
	}
	context.JSON(http.StatusOK, gin.H{"status": "success"})

	return nil
}

func UpdateEvent(event *Event, context *gin.Context) (*Event, error) {
	query := "UPDATE events SET name = ?, description = ?, location = ?, datetime = ?, userId = ? WHERE eventId = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId, &event.EventId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
