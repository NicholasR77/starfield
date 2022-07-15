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

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatal(err)
	}
	
	return &model.Ship{
		ID: res.InsertID.(primitive.ObjectID).Hex(),  
		Name: input.Name,
		Description: input.Description,
	}
}

funct (db * DB) FindByID(ID string) *model.Ship {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	
	if err != nil {
		log.Fatal(err)
	}

	shipsCollection := db.client.Database("starfield").Collection("ships")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne (ctx, bson.M{ "_id": ObjectID })
	ship := model.Ship{}
	res.Decode(&ship)

	return &ship
}

func (db *DB) All () []*model.Ship {
	shipsCollection := db.client.Database("starfield").Collection("ships")

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var ships []*model.Ships

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