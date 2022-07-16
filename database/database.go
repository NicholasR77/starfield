package database

import (
	"context"
	"log"
	"time"

	"github.com/NicholasR77/starfield/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}

func (db *DB) CreateShip(input *model.NewShip) *model.Ship {
	collection := db.client.Database("starfield").Collection("ships")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}

	return &model.Ship{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Name:        input.Name,
		Description: input.Description,
	}
}

// TODO: Look into why ID is not being returned properly
func (db *DB) CreateModule(ID string, input *model.NewModule) *model.Ship {
	ObjectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("starfield").Collection("ships")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$push": bson.M{"modules": input}}
	upsert := true
	after := options.After
	options := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	res := collection.FindOneAndUpdate(ctx, filter, update, &options)

	ship := model.Ship{}
	res.Decode(&ship)

	return &ship
}

func (db *DB) FindByID(ID string) *model.Ship {
	ObjectID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("starfield").Collection("ships")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	ship := model.Ship{}
	res.Decode(&ship)

	return &ship
}

func (db *DB) FindAll() []*model.Ship {
	collection := db.client.Database("starfield").Collection("ships")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var ships []*model.Ship

	for cur.Next(ctx) {
		var ship *model.Ship
		err := cur.Decode(&ship)

		if err != nil {
			log.Fatal(err)
		}

		ships = append(ships, ship)
	}

	return ships
}
