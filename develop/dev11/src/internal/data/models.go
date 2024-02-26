package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Events EventModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Events: EventModel{db: db},
	}
}
