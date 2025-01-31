package main

import (
	"fmt"
	"go-web-app/api"
	"go-web-app/components"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	flag "github.com/spf13/pflag"
)

var port = flag.IntP("port", "p", 3000, "Listen port")

func main() {
	flag.Parse()

	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339},
	).With().Timestamp().Logger()
	log.Logger = logger

	middleware.DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{Logger: &logger, NoColor: false},
	)

	if *port == 0 {
		log.Fatal().Msg("No --port was provided")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

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
	cs := components.NewService()

	r.Get("/", cs.Home)
	r.Mount("/components", cs.Routes())

	// API endpoints
	as := &api.Service{}
	r.Mount("/api", as.Routes())

	log.Info().Msg(fmt.Sprintf("Starting listening on :%d", *port))
	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}
