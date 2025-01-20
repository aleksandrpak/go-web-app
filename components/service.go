package components

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	count int
}

func (s *Service) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/counter", s.counter)

	return r
}

func (s *Service) counter(w http.ResponseWriter, r *http.Request) {
	s.count++
	counter(s.count).Render(r.Context(), w)
}
