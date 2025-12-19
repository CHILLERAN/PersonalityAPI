package main

import (
	"net/http"

	"github.com/CHILLERAN/PersonalityAPI/internal/config"
)

func notFoundError(app *config.Application, w http.ResponseWriter, r *http.Request, err error) {
	var(
		method = r.Method
		url = r.URL 
	)

	app.Logger.Error(err.Error(),"Method", method, "URL", url)
	http.NotFound(w, r)
}

func serverError(app *config.Application, w http.ResponseWriter, r *http.Request, err error) {
	var(
		method = r.Method
		url = r.URL 
	)

	app.Logger.Error(err.Error(),"Method", method, "URL", url)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}