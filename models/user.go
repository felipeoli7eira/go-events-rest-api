package models

import (
	"github.com/felipeoli7eira/go-events-rest-api/db"
)

type User struct {
	ID 		 int64
	Name 	 string `binding:"required"`
	email	 string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	stmt, err := db.Database.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Name, u.email, u.Password)

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
