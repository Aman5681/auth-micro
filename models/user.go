package models

import (
	"errors"

	"github.com/Aman5681/auth-micro/db"
	"github.com/Aman5681/auth-micro/utils"
)

type User struct {
	UserId   string `json:"userId"`
	EmailId  string `json:"emailId" bindings:"required"`
	Password string `json:"password" bindings:"required"`
}

func (u *User) Save() (userId *string, err error) {
	// Save user to database
	u.UserId = utils.GenerateUUID()
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

func (u *User) ValidateUser() error {
	// Validate user
	query := "SELECT userId, password FROM users WHERE emailId = ?"
	row := db.DB.QueryRow(query, u.EmailId)

	var retrievedPassword string
	err := row.Scan(&u.UserId, &retrievedPassword)

	if err != nil {
		return err
	}

	if !utils.ComparePassword(u.Password, retrievedPassword) {
		return errors.New("credentials invalid")
	}

	return nil
}
