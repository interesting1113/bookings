package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sylviayang/bookings/pkg/config"
	"github.com/sylviayang/bookings/pkg/handlers"
	"github.com/sylviayang/bookings/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {

		app.TemplateCache = tc
		app.UseCache = false

		repo := handlers.NewRepo(&app)
		handlers.NewHandler(repo)

		render.NewTemplate(&app)
		log.Fatal("cannot create template cache")
	}
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
