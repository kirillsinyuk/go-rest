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
	UserId      int64
}

func (event *Event) Save() error {
	insert := "INSERT INTO event (name, description, location, date_time, user_id) VALUES(?, ?, ?, ?, ?)"
	result, err := db.DB.Exec(insert, event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id

	return err
}

func (event *Event) Update() error {
	insert := "UPDATE event SET name = ?, description = ?, location = ?, date_time = ?, user_id = ? WHERE id = ?"
	_, err := db.DB.Exec(insert, event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.ID)

	return err
}

func DeleteById(ID int64) error {
	delete := "DELETE FROM event WHERE id = ?"
	_, err := db.DB.Exec(delete, ID)

	return err
}

func GetAllEvents() ([]Event, error) {
	selectAll := "SELECT * FROM event"
	rows, err := db.DB.Query(selectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := make([]Event, 0)
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(ID int64) (*Event, error) {
	row := db.DB.QueryRow("SELECT * FROM event WHERE id = ?", ID)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
