package models

import (
	"time"

	"github.com/felipeoli7eira/go-events-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var eventsDatabase []Event = []Event{}

func (event Event) SaveEvent() error {
	queryString := "INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.Database.Prepare(queryString)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)

	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	event.ID = lastInsertId

	return nil
}

func ListEvents() ([]Event, error) {
	rows, err := db.Database.Query("SELECT * FROM events")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	var queryString string = "SELECT * FROM events WHERE id = ?"

	row := db.Database.QueryRow(queryString, id)

	var e Event

	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}
