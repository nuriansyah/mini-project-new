package handlers

import (
	"github.com/go-chi/chi"
	"mini-project-new/internal/middleware"
	"net/http"
)

func RouteHandler(r chi.Router) http.Handler {
	r.Route("/v1/events", func(r chi.Router) {
		r.Group(func(r chi.Router) {

			r.Use(middleware.Authentication)

			r.Post("/", CreateEvents)

			r.Patch("/{events_id}", UpdateEvents)
			r.Delete("/{events_id}", DeleteEvents)
			r.Get("/{events_id}", GetEvent)

			r.Get("/", GetAllEvents)
		})
	})
	return r
}
