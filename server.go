package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/NicholasR77/starfield/graph"
	"github.com/NicholasR77/starfield/graph/generated"

	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

const graphqlPort = "8080"

func main() {
	// MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

    if err != nil {
        log.Fatal(err)
    }

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)

    if err != nil {
        log.Fatal(err)
    }

    defer client.Disconnect(ctx)
    
    databases, err := client.ListDatabaseNames(ctx, bson.M{})

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(databases)

	// GraphQL
	port := os.Getenv("PORT")

	if port == "" {
		port = graphqlPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
