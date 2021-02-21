package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/madegun/bookings/pkg/config"
	"github.com/madegun/bookings/pkg/handlers"
	"github.com/madegun/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const port = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

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

	fmt.Println(fmt.Sprintf("server running di %s", port))

	serv := &http.Server{
		Addr:    port,
		Handler: Routes(&app),
	}
	err = serv.ListenAndServe()
	log.Fatal(err)
}
