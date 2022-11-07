package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/*
NoSurf adds CSRF protection to all POST requests

	Includes peramiters for use using the nosurf package
*/
func NoSurf(next http.Handler) http.Handler {
	//Create a new token
	csrfHandler := nosurf.New(next)

	//cookie peramiters
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessonLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
