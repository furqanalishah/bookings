package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/furqanalishah/bookings/pkg/config"
	"github.com/furqanalishah/bookings/pkg/handlers"
	"github.com/furqanalishah/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to `true` when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// set the app.Session
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Println(fmt.Sprintf("Starting web server on the http://localhost%s", PORT))

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("error starting the server:", err)
	}
}
