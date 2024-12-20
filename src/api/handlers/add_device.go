package handlers

import (
	"encoding/json"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/requests"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func AddDevice(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDevice(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lastSyncTime := primitive.Timestamp{
		T: uint32(req.LastSyncTime),
		I: 0,
	}

	device := data.Device{
		ID:           primitive.NewObjectID(),
		PetID:        req.PetID,
		Status:       req.Status,
		LastSyncTime: lastSyncTime,
	}

	devicesDB := MongoDB(r).Devices()

	err = devicesDB.Insert(&device)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key error collection") {
			w.WriteHeader(http.StatusConflict)
			return
		}
		http.Error(w, "Failed to add device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Device added successfully",
		"deviceID": device.ID.Hex(),
	})
}
