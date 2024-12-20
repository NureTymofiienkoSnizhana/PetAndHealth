package requests

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

type UpdateUser struct {
	ID       primitive.ObjectID   `json:"_id"`
	FullName string               `json:"full_name"`
	Role     string               `json:"role"`
	Email    string               `json:"email"`
	Password string               `json:"password"`
	PetsID   []primitive.ObjectID `json:"pets_id"`
}

func NewUpdateUser(r *http.Request) (*UpdateUser, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var user UpdateUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
