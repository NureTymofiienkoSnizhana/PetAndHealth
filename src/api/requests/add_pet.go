package requests

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

type AddPet struct {
	Name    string             `json:"name"`
	Species string             `json:"species"`
	Breed   string             `json:"breed"`
	Age     int                `json:"age"`
	OwnerID primitive.ObjectID `json:"owner_id"`
}

func NewPet(r *http.Request) (*AddPet, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var pet AddPet
	err = json.Unmarshal(body, &pet)
	if err != nil {
		return nil, err
	}

	return &pet, nil
}
