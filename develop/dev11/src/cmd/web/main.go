package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"time"

	"dev11/src/internal/config"
	"dev11/src/internal/data"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	models data.Models
}

func main() {
	cfg := config.MustLoad(".env")

	db, err := sql.Open("sqlite3", cfg.DBPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		models: data.NewModels(db),
	}

	err = app.models.Events.InitDB(db)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         cfg.Address(),
		Handler:      logMiddleware(app.routes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	slog.Info("Server is listening...", "address", cfg.Address())

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
