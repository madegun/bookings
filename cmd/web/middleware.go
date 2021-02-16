package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf middleware to prevent from Cross-Site Request Forgery attacks
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//SessionLoad middleware to load and save every request page
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
