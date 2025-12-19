package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/CHILLERAN/PersonalityAPI/internal/config"
	"github.com/CHILLERAN/PersonalityAPI/internal/models"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	loggerhandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerhandler)

	db, err := sql.Open("mysql", "web:password@/personalityapi")

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := config.Application{
		TraitModel: &models.TraitModel{DB: db},
		Logger: logger,
	}

	logger.Info("Starting Server", "Host", "http://localhost:4000")

	err = http.ListenAndServe(":4000", routes(&app))

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func OpenDB(driver, connectionString string) (*sql.DB, error) {
	db, err := sql.Open(driver, connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}