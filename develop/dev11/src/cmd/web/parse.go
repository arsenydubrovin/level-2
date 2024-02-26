package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"dev11/src/internal/data"
)

func (app *application) parseDateParam(r *http.Request) (time.Time, error) {
	param := r.URL.Query().Get("date")
	if param == "" {
		return time.Time{}, errors.New("date parameter is required")
	}

	date, err := time.Parse(time.DateOnly, param)
	if err != nil {
		return time.Time{}, errors.New("invalid date")
	}

	return date, nil
}

func (app *application) parseIdParam(r *http.Request) (int64, error) {
	err := r.ParseForm()
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

func (app *application) parseEventForm(r *http.Request, event *data.Event) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	if titleParam := r.Form.Get("title"); titleParam != "" {
		event.Title = titleParam
	}

	if descParam := r.Form.Get("desc"); descParam != "" {
		event.Desc = descParam
	}

	if startsAtParam := r.Form.Get("starts_at"); startsAtParam != "" {
		startsAt, err := time.Parse(time.RFC3339, startsAtParam)
		if err != nil {
			return errors.New("invalid starts_at")
		}
		event.StartsAt = startsAt
	}

	if endsAtParam := r.Form.Get("ends_at"); endsAtParam != "" {
		endsAt, err := time.Parse(time.RFC3339, endsAtParam)
		if err != nil {
			return errors.New("invalid ends_at")
		}
		event.EndsAt = endsAt
	}

	return nil
}
