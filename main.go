package main

import (
	"context"
	"fmt"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func main() {
	ctx := context.Background()

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	if username == "" || password == "" {
		panic("MONGO_USERNAME або MONGO_PASSWORD не встановлені")
	}

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@petandhealth.dtpxu.mongodb.net/?retryWrites=true&w=majority&appName=PetAndHealth", username, password)
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	petAndHealthDB := client.Database("PetAndHealth")
	mongoDB := mongodb.NewMasterDB(petAndHealthDB)

	api.Run(api.Config{
		MasterDB: mongoDB,
	})
}
