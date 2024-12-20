package mongodb

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const NotificationsCollectionName = "Notifications"

type notificationsDB struct {
	collection *mongo.Collection
}

func newNotificationsDB(db *mongo.Database) *notificationsDB {
	return &notificationsDB{
		collection: db.Collection(NotificationsCollectionName),
	}
}

func NewNotificationsDB(db *mongo.Database) data.NotificationsDB {
	return newNotificationsDB(db)
}

func (n *notificationsDB) Get(id string) (*data.Notification, error) {
	var result data.Notification
	err := n.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (n *notificationsDB) Insert(notif *data.Notification) error {
	_, err := n.collection.InsertOne(context.TODO(), notif)
	return err
}

func (n *notificationsDB) Update(id string, updateFields bson.M) error {
	_, err := n.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updateFields},
	)
	return err
}
