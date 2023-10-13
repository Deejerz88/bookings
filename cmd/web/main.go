package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Deejerz88/bookings/pkg/config"
	"github.com/Deejerz88/bookings/pkg/handlers"
	"github.com/Deejerz88/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const protNumber = ":8080"

var app = config.AppConfig{}
var session *scs.SessionManager

func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Error creating template cache :", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Server is running on port %s", protNumber)

	srv := &http.Server{
		Addr:    protNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
