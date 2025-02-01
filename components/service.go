package components

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Service struct {
	count int
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/counter", s.counter)

	return r
}

func (s *Service) Home(w http.ResponseWriter, r *http.Request) {
	home().Render(r.Context(), w)
}

func (s *Service) counter(w http.ResponseWriter, r *http.Request) {
	s.count++

	if s.count > 3 {
		log.Error().Err(fmt.Errorf("Unknown error")).Msg("Encountered error")

		w.WriteHeader(http.StatusBadRequest)
		errorToast("can't count anymore").Render(r.Context(), w)
	} else {
		counter(s.count).Render(r.Context(), w)
	}
}
