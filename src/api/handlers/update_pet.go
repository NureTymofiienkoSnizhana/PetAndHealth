package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdatePet(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateFields := bson.M{
		"name":     req.Name,
		"species":  req.Species,
		"breed":    req.Breed,
		"age":      req.Age,
		"owner_id": req.OwnerID,
	}

	if req.ID.IsZero() {
		http.Error(w, "Invalid pet ID", http.StatusBadRequest)
		return
	}

	petsDB := MongoDB(r).Pets()

	err = petsDB.Update(req.ID, updateFields)
	if err != nil {
		http.Error(w, "Failed to update pet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Pet updated successfully",
		"petID":   req.ID.Hex(),
	})
}
