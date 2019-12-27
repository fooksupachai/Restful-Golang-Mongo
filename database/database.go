package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Member struct
type Member struct {
	A int `json:"a"`
}

// InitialDB function
func InitialDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:2277")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// GetAllMember function
func GetAllMember(client *mongo.Client, filter bson.M) []*Member {
	var heroes []*Member
	collection := client.Database("pda").Collection("documents")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero Member
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}

// GetMember function
func GetMember() {

}
