package requests

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

type UpdateDevice struct {
	ID           primitive.ObjectID `json:"_id"`
	PetID        primitive.ObjectID `json:"pet_id"`
	Status       string             `json:"status"`
	LastSyncTime int64              `json:"last_sync_time"`
}

func NewUpdateDevice(r *http.Request) (*UpdateDevice, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var device UpdateDevice
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}
