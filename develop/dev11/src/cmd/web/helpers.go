package main

import (
	"time"
)

type timeBounds struct {
	start time.Time
	end   time.Time
}

func (app *application) getDayBounds(t time.Time) timeBounds {
	startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Second)

	return timeBounds{
		start: startOfDay,
		end:   endOfDay,
	}
}

func (app *application) getWeekBounds(t time.Time) timeBounds {
	daysUntilMonday := (int(t.Weekday()) + 6) % 7

	startOfWeek := time.Date(t.Year(), t.Month(), t.Day()-daysUntilMonday, 0, 0, 0, 0, t.Location())
	endOfWeek := startOfWeek.Add(7 * 24 * time.Hour).Add(-time.Second)

	return timeBounds{
		start: startOfWeek,
		end:   endOfWeek,
	}
}

func (app *application) getMonthBounds(t time.Time) timeBounds {
	startOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	return timeBounds{
		start: startOfMonth,
		end:   endOfMonth,
	}
}
