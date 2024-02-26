package data

import (
	"database/sql"
	"time"

	"dev11/src/internal/validator"
)

type Event struct {
	Id       int64     `json:"id"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
	Title    string    `json:"title"`
	Desc     string    `json:"desc"`
}

func (a *Event) Validate(v *validator.Validator) {
	v.Check(!a.StartsAt.IsZero(), "starts_at", "must be provided")
	v.Check(!a.EndsAt.IsZero(), "ends_at", "must be provided")

	v.Check(a.StartsAt.Before(a.EndsAt) || a.StartsAt.Equal(a.EndsAt), "starts_at", "must be less than or equal to ends_at")

	v.Check(a.Title != "", "title", "must be provided")
	v.Check(len([]rune(a.Title)) <= 200, "title", "must be no more than 200 characters")

	// description can be empty
	v.Check(len([]rune(a.Desc)) <= 1000, "description", "must be no more than 1000 characters")
}

type EventModel struct {
	db *sql.DB
}

func (em *EventModel) InitDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS events (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						starts_at DATETIME,
						ends_at DATETIME,
						title TEXT,
						desc TEXT)`

	_, err := db.Exec(query)
	return err
}

func (em *EventModel) Get(id int64) (*Event, error) {
	query := `SELECT id, starts_at, ends_at, title, desc
						FROM events
						WHERE id = ?`

	var event Event

	err := em.db.QueryRow(query, id).Scan(
		&event.Id,
		&event.StartsAt,
		&event.EndsAt,
		&event.Title,
		&event.Desc)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (em *EventModel) GetByTime(from, to time.Time) ([]Event, error) {
	query := `SELECT id, starts_at, ends_at, title, desc
						FROM events
						WHERE starts_at >= ? AND starts_at <= ?`

	rows, err := em.db.Query(query, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.Id,
			&event.StartsAt,
			&event.EndsAt,
			&event.Title,
			&event.Desc)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (em *EventModel) Insert(event *Event) (int64, error) {
	query := `INSERT INTO events (starts_at, ends_at, title, desc)
						VALUES (?, ?, ?, ?)`

	result, err := em.db.Exec(
		query,
		event.StartsAt,
		event.EndsAt,
		event.Title,
		event.Desc)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (em *EventModel) Update(event *Event) error {
	query := `UPDATE events
						SET starts_at=?, ends_at=?, title=?, desc=?
						WHERE id=?`

	result, err := em.db.Exec(
		query,
		event.StartsAt,
		event.EndsAt,
		event.Title,
		event.Desc,
		event.Id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (em *EventModel) Delete(id int64) error {
	query := `DELETE FROM events WHERE id=?`

	result, err := em.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
