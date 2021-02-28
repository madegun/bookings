package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/madegun/bookings/internal/config"
	"github.com/madegun/bookings/internal/handlers"
	"github.com/madegun/bookings/internal/render"
	"github.com/madegun/bookings/models"

	"github.com/alexedwards/scs/v2"
)

const port = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//apa yg akan disimpan di session
	gob.Register(models.Reservation{})

	//true/false production mode
	app.InProduction = false

	//variable new session dan settingannya menggunakan middleware scs
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateChace()
	if err != nil {
		log.Fatal("tidak bisa render template ke browser")
	}

	app.TemplateChace = tc
	app.UseChace = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("server running port %s", port))

	serv := &http.Server{
		Addr:    port,
		Handler: Routes(&app),
	}
	err = serv.ListenAndServe()
	log.Fatal(err)
}
