package database

import {
	"context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/NicholasR77/starfield/graph/model"
}

type DB struct [
	client *mongo.Client 
]

func Connect() *DB  {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	return &DB { 
		client: client,
	}
}

func (db* DB) Save(input *model.NewShip) *model.Ship {
	shipsCollection := db.client.Database("starfield").Collection("ships")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if error != nil {
		log.Fatal(err)
	}

	res, err := collection.InsertOne(ctx, input)
	
	return &model.Ship{
		ID: res.InsertID.(primitive.ObjectID).Hex(),
		Name: input.Name,
		Description: input.Description,
	}
}