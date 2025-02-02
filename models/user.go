package models

import (
	"errors"

	"github.com/felipeoli7eira/go-events-rest-api/utils"
	"github.com/felipeoli7eira/go-events-rest-api/db"
)

type User struct {
	ID 		 int64
	Name 	 string `binding:"required"`
	Email	 string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	stmt, err := db.Database.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPlainText(u.Password)

	if err != nil {
		return err
	}

	u.Password = hashedPassword

	result, err := stmt.Exec(u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = lastInsertId

	return nil
}

func (u User) ValidateCredentials() error {
	row := db.Database.QueryRow("SELECT password FROM users WHERE email = ?", u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordMatch := utils.PlainTextAndHashMatch(retrievedPassword, u.Password)

	if !passwordMatch {
		return errors.New("Invalid credentials")
	}

	return nil
}
