package api

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/handlers"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/middle"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"time"
)

type Config struct {
	MasterDB data.MasterDB
}

func Run(config Config) {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), handlers.MasterDBContextKey, config.MasterDB)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
	r.Use(middleware.Timeout(60 * time.Second))

	// Routes
	r.Route("/api/pet-and-health", func(r chi.Router) {
		r.Route("/v1/login", func(r chi.Router) {
			r.Get("/auth", handlers.Auth)
			r.Post("/registration", handlers.Registration)
			r.Post("/create-admin", handlers.CreateAdmin)
		})
		r.Route("/v1/admin/pets", func(r chi.Router) {
			r.Use(middle.MockUser("user"))
			r.Use(middle.CheckRole("admin"))
			r.Post("/add-pet", handlers.AddPet)
			r.Delete("/delete-pet", handlers.DeletePet)
			r.Get("/get-pets", handlers.GetPets)
			r.Put("/update-pet", handlers.UpdatePet)
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
