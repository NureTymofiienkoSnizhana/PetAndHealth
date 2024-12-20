package requests

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

type AddDevice struct {
	PetID        primitive.ObjectID `json:"pet_id"`
	Status       string             `json:"status"`
	LastSyncTime int64              `json:"last_sync_time"`
}

func NewDevice(r *http.Request) (*AddDevice, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var device AddDevice
	err = json.Unmarshal(body, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}
