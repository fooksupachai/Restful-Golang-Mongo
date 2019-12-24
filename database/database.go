package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Member struct
type Member struct {
	_id string
	a   int
}

// InitialDB function
func InitialDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:2277")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// GetMember function
func GetMember() bson.M {
	collection := InitialDB().Database("pda").Collection("documents")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result bson.M
	for cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	InitialDB().Disconnect(context.TODO())
	fmt.Println("Connection to MongoDB closed.")

	return result
}
