package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NotificationsDB interface {
	Get(id string) (*Notification, error)
	Insert(*Notification) error
	Update(id string, updateFields bson.M) error
}

type Notification struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  primitive.ObjectID `bson:"user_id"`
	Message string             `bson:"message"`
	Time    time.Time          `bson:"time"`
}
