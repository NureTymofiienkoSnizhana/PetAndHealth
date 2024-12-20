package requests

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
)

type PetReport struct {
	PetID     primitive.ObjectID `json:"_id"`
	StartTime int64              `json:"start_time"`
	EndTime   int64              `json:"end_time"`
}

func NewPetReport(r *http.Request) (*PetReport, error) {
	bodyReader := r.Body
	if bodyReader == nil {
		return nil, errors.New("missing body")
	}

	body, err := io.ReadAll(bodyReader)
	if err != nil {
		return nil, err
	}

	var report PetReport
	err = json.Unmarshal(body, &report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
