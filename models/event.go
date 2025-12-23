package models

import (
	"rest/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (event Event) Save() error {
	insert := `INSERT INTO events (id, name, description, location, date_time, user_id) VALUES(?, ?, ?, ?, ?, ?)`
	result, err := db.DB.Exec(insert, event.ID, event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	selectAll := `SELECT * FROM events`
	rows, err := db.DB.Query(selectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := make([]Event, 0)
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(ID int64) (*Event, error) {
	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", ID)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
