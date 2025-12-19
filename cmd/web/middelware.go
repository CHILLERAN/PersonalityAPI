package main

import (
	"fmt"
	"net/http"

	"github.com/CHILLERAN/PersonalityAPI/internal/config"
)

func commonHeaders(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
		"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		w.Header().Set("Server", "Go")

		next.ServeHTTP(w,r)
	}
}

func logRequests(app *config.Application, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var(
			method = r.Method
			url = r.URL 
			proto = r.Proto
		)

		app.Logger.Info("Request","Method", method, "URL", url, "Protocol", proto)

		next.ServeHTTP(w,r)
	}
}

func panicRecovery(app *config.Application, next http.Handler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			panicHappend := recover()

			if panicHappend != nil {
				w.Header().Set("Connection", "close")

				serverError(app, w, r, fmt.Errorf("%v", panicHappend))
			}
		}()

		next.ServeHTTP(w, r)
	}
}