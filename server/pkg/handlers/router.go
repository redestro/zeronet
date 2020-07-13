package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

//InitRoutes exports the router
func InitRoutes(client *mongo.Client) *chi.Mux {
	db := client.Database("zeronet")

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
		r.Get("/human", func(w http.ResponseWriter, r *http.Request) {
			StartHuman(w, r, db)
		})
		r.Get("/ai", func(w http.ResponseWriter, r *http.Request) {
			StartAI(w, r, db)
		})
	})
	return r
}
