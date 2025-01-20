package main

import (
	"net/http"
	"templateapp/api"
	"templateapp/components"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Main page
	r.Get("/", templ.Handler(components.Home()).ServeHTTP)

	// Static content
	r.Handle(
		"/assets/js/*",
		http.StripPrefix("/assets/js", http.FileServer(http.Dir("./assets/js"))),
	)
	r.Handle(
		"/assets/css/*",
		http.StripPrefix("/assets/css", http.FileServer(http.Dir("./assets/css"))),
	)

	// HTMX components
	cs := &components.Service{}
	r.Mount("/components", cs.Routes())

	// API endpoints
	as := &api.Service{}
	r.Mount("/api", as.Routes())

	http.ListenAndServe(":3000", r)
}
