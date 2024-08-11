package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Maged-Zaki/gin-rest-api/db"
)

type Event struct {
	ID          int64     `json:"id,omitempty"` // omitempty means if field has its null/default value then it's not returned in json
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserID      int64     `json:"userId"`
}

func (event *Event) Save() error {
	query := "INSERT INTO events(name, description, location, date, user_id) VALUES(?, ?, ?, ?, ?) RETURNING id"

	// Execute query and scan the returned id
	err := db.DB.QueryRow(query, event.Name, event.Description, event.Location, event.Date, event.UserID).Scan(&event.ID)
	if err != nil {
		return err
	}

	return nil
}

func (event *Event) Update() error {
	query := "UPDATE events SET name=?, description=?, location=?, date=? WHERE id=?"

	// Execute query
	result, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.Date, event.ID)
	if err != nil {
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, return a custom error
	if rowsAffected == 0 {
		return fmt.Errorf("no event found with the specified ID")
	}

	return nil
}
func (event *Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"

	// Execute query
	result, err := db.DB.Exec(query, event.ID)
	if err != nil {
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, return a custom error
	if rowsAffected == 0 {
		return fmt.Errorf("no event found with the specified ID")
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	// populate events of Event type
	var events []Event

	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		events = append(events, event)
	}

	return events, nil
}
func GetEvent(id int64) (Event, error) {
	query := "SELECT * FROM events WHERE id = ?"

	var e Event

	row := db.DB.QueryRow(query, id)
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.Date, &e.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Event{}, fmt.Errorf("Event with ID %d not found", id)
		}
		return Event{}, err
	}

	return e, nil
}
