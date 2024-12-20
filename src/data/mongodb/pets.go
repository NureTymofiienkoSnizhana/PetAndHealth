package mongodb

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const PetsCollectionName = "Pets"

type petsDB struct {
	collection *mongo.Collection
}

func newPetsDB(db *mongo.Database) *petsDB {
	return &petsDB{
		collection: db.Collection(PetsCollectionName),
	}
}

func NewPetsDB(db *mongo.Database) data.PetsDB {
	return newPetsDB(db)
}

func (p *petsDB) Get(id primitive.ObjectID) (*data.Pet, error) {
	var result data.Pet
	err := p.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *petsDB) Insert(pet *data.Pet) error {
	_, err := p.collection.InsertOne(context.TODO(), pet)
	return err
}

func (p *petsDB) Update(id primitive.ObjectID, updateFields bson.M) error {
	_, err := p.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$set": updateFields},
	)
	return err
}

func (p *petsDB) Delete(id primitive.ObjectID) error {
	_, err := p.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (p *petsDB) GetAll() ([]*data.Pet, error) {
	var pets []*data.Pet
	cursor, err := p.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var pet *data.Pet
		err := cursor.Decode(&pet)
		if err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}

func (p *petsDB) GetPetWithHealth(petID primitive.ObjectID) (*data.PetWithHealth, error) {
	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{{"_id", petID}}},
		},
		{
			{"$lookup", bson.D{
				{"from", "HealthData"},
				{"localField", "_id"},
				{"foreignField", "pet_id"},
				{"as", "health_data"},
			}},
		},
	}

	cursor, err := p.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []data.PetWithHealth
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &results[0], nil
}

func (p *petsDB) GetByFilter(filter bson.M) ([]*data.Pet, error) {
	var pets []*data.Pet
	cursor, err := p.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var pet data.Pet
		if err := cursor.Decode(&pet); err != nil {
			return nil, err
		}
		pets = append(pets, &pet)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return pets, nil
}
