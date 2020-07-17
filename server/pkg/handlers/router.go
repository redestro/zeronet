package handlers

import (
	"net/http"
	"time"

	"github.com/redestro/zeronet/server/pkg/middlewares"
	"github.com/redestro/zeronet/server/pkg/sessions"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

//InitRoutes exports the router
func InitRoutes(db *sessions.Database) *chi.Mux {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "x-access-token"},
		ExposedHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	})

	// middlewares
	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You have reached zeronet"))
	})

	r.Route("/start", func(r chi.Router) {
		r.Get("/human/{level1}/{level2}", StartHuman)
		r.Get("/ai/{level1}/{level2}", StartAI)
		r.Get("/revhuman/{level1}/{level2}", StartRevHuman)
		r.Get("/human2/{level1}/{level2}", StartHuman2)
	})
	r.Route("/play/{token}", func(r chi.Router) {
		r.Use(middlewares.Update)
		r.Post("/human", PlayHuman)
		r.Post("/ai", PlayAI)
	})
	r.Delete("/session/{token}", DeleteSession)
	return r
}
