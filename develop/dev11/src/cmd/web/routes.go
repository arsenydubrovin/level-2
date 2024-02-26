package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/events_for_day", app.getEventsForDayHandler)
	mux.HandleFunc("/events_for_week", app.getEventsForWeekHandler)
	mux.HandleFunc("/events_for_month", app.getEventsForMonthHandler)

	mux.HandleFunc("/create_event", app.createEventHandler)
	mux.HandleFunc("/update_event", app.updateEventHandler)
	mux.HandleFunc("/delete_event", app.deleteEventHadler)

	return mux
}
