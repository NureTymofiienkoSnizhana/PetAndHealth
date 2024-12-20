package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevicesDB interface {
	Get(id primitive.ObjectID) (*Device, error)
	Insert(*Device) error
	Update(id primitive.ObjectID, updateFields bson.M) error
}

type Device struct {
	ID           primitive.ObjectID  `bson:"_id"`
	PetID        primitive.ObjectID  `bson:"pet_id"`
	Status       string              `bson:"status"`
	LastSyncTime primitive.Timestamp `bson:"last_sync_time"`
}
