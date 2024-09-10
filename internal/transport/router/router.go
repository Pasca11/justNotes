package router

import (
	"github.com/Pasca11/justNotes/internal/transport/controller"
	"github.com/go-chi/chi/v5"
)

func New(c controller.Controller) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", c.HelloWorld)

	return r
}
