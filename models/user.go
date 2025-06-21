package models

import (
	"github.com/Aman5681/auth-micro/db"
	"github.com/Aman5681/auth-micro/utils"
	"github.com/google/uuid"
)

type User struct {
	UserId   uuid.UUID `json:"userId"`
	EmailId  string    `json:"emailId" bindings:"required"`
	Password string    `json:"password" bindings:"required"`
}

func (u *User) Save() (userId *uuid.UUID, err error) {
	// Save user to database
	u.UserId = uuid.New()
	query := "INSERT INTO users(userId, emailId, password) VALUES (?, ?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(u.UserId, u.EmailId, hashedPassword)

	if err != nil {
		return nil, err
	}

	return &u.UserId, nil
}
