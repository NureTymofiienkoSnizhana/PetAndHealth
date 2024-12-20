package api

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/handlers"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/middle"
	"github.com/go-chi/chi/v5"

	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type Config struct {
	MasterDB data.MasterDB
}

func Run(config Config) {
	r := chi.NewRouter()

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

	r.Route("/api/pet-and-health", func(r chi.Router) {
		r.Route("/v1/login", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Get("/auth", handlers.Auth)
			r.Post("/registration", handlers.Registration)
			r.Post("/create-admin", handlers.CreateAdmin)
		})
		r.Route("/v1/admin/pets", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Use(middle.MockUser("user"))
			r.Use(middle.CheckRole("admin"))
			r.Post("/add-pet", handlers.AddPet)
			r.Delete("/delete-pet", handlers.DeletePet)
			r.Get("/get-pets", handlers.GetPets)
			r.Put("/update-pet", handlers.UpdatePet)
		})
		r.Route("/v1/admin/devices", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Use(middle.MockUser("admin"))
			r.Use(middle.CheckRole("admin"))
			r.Post("/add-device", handlers.AddDevice)
			r.Put("/update-device", handlers.UpdateDevice)
		})
		r.Route("/v1/admin/users", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Use(middle.MockUser("admin"))
			r.Use(middle.CheckRole("admin"))
			r.Put("/update-user", handlers.UpdateUser)
			r.Delete("/delete-user", handlers.DeleteUser)
		})
		r.Route("/v1/user:id", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Put("/update-user", handlers.UpdateUser)
			r.Get("/profile", handlers.UserInfo)
			//r.Post("/logout", handlers.Logout)
		})
		r.Route("/v1/user", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Get("/pet-info", handlers.PetInfo)
		})
		r.Route("/v1/owner", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Use(middle.MockUser("user"))
			r.Use(middle.CheckRole("user"))
			r.Get("/owner-pets", handlers.GetOwnerPets)
		})
		r.Route("/v1/vet", func(r chi.Router) {
			//r.Use(middle.AuthMiddleware)
			r.Use(middle.MockUser("user"))
			r.Use(middle.CheckRole("vet"))
			r.Get("/pet-report/pdf", handlers.GetPetReport)
			r.Get("/get-pets", handlers.GetPets)
		})
		r.Route("/v1/health-data", func(r chi.Router) {
			r.Post("/device", handlers.AddHealthData)
		})
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
