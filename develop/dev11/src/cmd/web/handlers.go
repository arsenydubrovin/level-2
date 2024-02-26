package main

import (
	"errors"
	"net/http"
	"time"

	"dev11/src/internal/data"
	"dev11/src/internal/validator"
)

func (app *application) getEventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	app.getEventsForBounds(w, r, app.getDayBounds)
}

func (app *application) getEventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	app.getEventsForBounds(w, r, app.getWeekBounds)
}

func (app *application) getEventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	app.getEventsForBounds(w, r, app.getMonthBounds)
}

func (app *application) getEventsForBounds(w http.ResponseWriter, r *http.Request, getTimeBounds func(time.Time) timeBounds) {
	if r.Method != http.MethodGet {
		app.methodNotAllowedResponse(w, r)
		return
	}

	date, err := app.parseDateParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	tb := getTimeBounds(date)
	events, err := app.models.Events.GetByTime(tb.start, tb.end)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"events": events, "count": len(events)}, nil)
}

func (app *application) createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.methodNotAllowedResponse(w, r)
		return
	}

	event := &data.Event{}
	err := app.parseEventForm(r, event)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	v := validator.New()

	event.Validate(v)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	id, err := app.models.Events.Insert(event)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"result": "event created successfully", "id": id}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.methodNotAllowedResponse(w, r)
		return
	}

	id, err := app.parseIdParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	event, err := app.models.Events.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.parseEventForm(r, event)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	event.Validate(v)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Events.Update(event)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"result": "event updated successfully", "id": id}, nil)
}

func (app *application) deleteEventHadler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.methodNotAllowedResponse(w, r)
		return
	}

	id, err := app.parseIdParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Events.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"result": "event deleted successfully"}, nil)
}
