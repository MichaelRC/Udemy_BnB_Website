package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MichaelRC/Udemy_BnB_Website/pkg/config"
	"github.com/MichaelRC/Udemy_BnB_Website/pkg/handlers"
	"github.com/MichaelRC/Udemy_BnB_Website/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	app.InProduction = false //<---!!!change this to true when in production

	//starting a new session
	session = scs.New()
	//how long the session will last (24 hours)
	session.Lifetime = 24 * time.Hour
	//session will persist even when browser is closed
	session.Cookie.Persist = true
	//set how strict which site this cookie applies to
	session.Cookie.SameSite = http.SameSiteLaxMode
	//this is to make sure cookies are secured with HTTPS
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannont create template cache.")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//give our render access to app config
	//using '&' as a referance to a pointer
	render.NewTemplates(&app)

	//prints to the console to make notify that program is running.
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
