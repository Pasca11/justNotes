package router

import (
	"github.com/Pasca11/justNotes/internal/transport/controller"
	"github.com/Pasca11/justNotes/internal/transport/controller/mw"
	"github.com/go-chi/chi/v5"
)

func New(c controller.Controller) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", c.Login)
		r.Post("/register", c.Register)
	})
	r.Route("/api", func(r chi.Router) {
		r.Use(mw.AuthenticationMiddleware)

		r.Get("/notes", c.GetNotes)
		r.Post("/notes", c.CreateNote)
		r.Delete("/notes", c.DeleteNote)

	})
	return r
}
