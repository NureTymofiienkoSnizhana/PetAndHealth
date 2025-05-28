package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data/mongodb"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Отриманий запит: %s %s", r.Method, r.URL.Path)

	godotenv.Load()

	originalPath := r.URL.Path

	if strings.HasPrefix(originalPath, "/api") {
		r.URL.Path = strings.TrimPrefix(originalPath, "/api")
	}

	if r.URL.Path == "" {
		r.URL.Path = "/"
	}

	r.URL.Path = "/api/pet-and-health" + r.URL.Path

	log.Printf("Змінений шлях: %s", r.URL.Path)

	ctx := context.Background()

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	if username == "" || password == "" {
		log.Printf("Помилка: відсутні дані MongoDB")
		http.Error(w, "MONGO_USERNAME або MONGO_PASSWORD не встановлені", http.StatusInternalServerError)
		return
	}

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@petandhealth.dtpxu.mongodb.net/?retryWrites=true&w=majority&appName=PetAndHealth", username, password)
	clientOptions := options.Client().ApplyURI(mongoURI)

	clientOptions.SetConnectTimeout(10 * time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("Помилка підключення до MongoDB: %v", err)
		http.Error(w, "Помилка підключення до MongoDB: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Не вдалось підключитись до MongoDB: %v", err)
		http.Error(w, "Не вдалось встановити зв'язок з MongoDB: "+err.Error(), http.StatusInternalServerError)
		return
	}

	petAndHealthDB := client.Database("PetAndHealth")
	mongoDB := mongodb.NewMasterDB(petAndHealthDB)

	router := api.GetRouter(api.Config{
		MasterDB: mongoDB,
	})

	log.Printf("Передаємо запит маршрутизатору: %s %s", r.Method, r.URL.Path)
	router.ServeHTTP(w, r)
}
