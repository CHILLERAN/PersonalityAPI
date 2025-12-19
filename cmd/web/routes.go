package main

import (
	"net/http"

	"github.com/CHILLERAN/PersonalityAPI/internal/config"
)

func routes(app *config.Application) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /{$}", home(app))
	mux.Handle("GET /category/{category}", getAllTraitsByCategory(app))
	mux.Handle("GET /category/{category}/trait/{trait}", getUniqueTraitByCategory(app))

	return panicRecovery(app, logRequests(app, commonHeaders(mux)))
}