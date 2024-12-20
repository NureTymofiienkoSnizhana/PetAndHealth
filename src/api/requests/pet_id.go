package requests

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

type PetID struct {
	ID primitive.ObjectID `json:"_id"`
}

func NewPetID(r *http.Request) (*PetID, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var pet PetID
	err = json.Unmarshal(body, &pet)
	if err != nil {
		return nil, err
	}

	return &pet, nil
}
