package router

import (
	_ "github.com/Pasca11/justNotes/docs"
	"github.com/Pasca11/justNotes/internal/transport/controller"
	"github.com/Pasca11/justNotes/internal/transport/router/mw"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func New(c controller.Controller) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", c.Login)
		r.Post("/register", c.Register)
	})
	r.Route("/api", func(r chi.Router) {
		r.Use(mw.LatMetricsMiddleware)
		r.Use(mw.AuthenticationMiddleware)
		r.Use(mw.AdminOnlyMiddleware)

		r.Get("/{id}/notes", c.GetNotes)
		r.Post("/{id}/notes", c.CreateNote)
		r.Delete("/notes/{id}", c.DeleteNote)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	r.Handle("/metrics", promhttp.Handler())
	return r
}
