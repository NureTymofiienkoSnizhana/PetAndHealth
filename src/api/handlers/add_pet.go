package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func AddPet(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewPet(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pet := data.Pet{
		ID:      primitive.NewObjectID(),
		Name:    req.Name,
		Species: req.Species,
		Breed:   req.Breed,
		Age:     req.Age,
		OwnerID: req.OwnerID,
	}

	petsDB := MongoDB(r).Pets()

	err = petsDB.Insert(&pet)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error collection") {
			w.WriteHeader(http.StatusConflict)
			return
		}
		http.Error(w, "Failed to add pet", http.StatusInternalServerError)
		return
	}

	usersDB := MongoDB(r).Users()

	err = usersDB.UpdatePets(req.OwnerID, pet.ID)
	if err != nil {
		http.Error(w, "Failed to update user with pet ID", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Pet added successfully",
		"petID":   pet.ID.Hex(),
	})
}
