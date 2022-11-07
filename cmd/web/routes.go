package main

import (
	"net/http"

	"github.com/MichaelRC/Udemy_BnB_Website/pkg/config"
	"github.com/MichaelRC/Udemy_BnB_Website/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes() uses the chi router for Golang HTTP services
func routes(app *config.AppConfig) http.Handler {
	//create NewRouter and store in variable mux
	mux := chi.NewRouter()

	/* Add middleware */

	//Revoverer restart program after panic and logs info
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	/*mux.Get retreaves the template of the page to use */
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//return what is stored in mux
	return mux
}
