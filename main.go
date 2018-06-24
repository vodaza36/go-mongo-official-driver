package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

const (
	mongoURL = "mongodb://localhost:27017"
)

// User entity
type User struct {
	ID    string `bson:"id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func createDBConnection() *mongo.Database {
	client, err := mongo.Connect(context.Background(), mongoURL, nil)

	if err != nil {
		panic(err)
	}

	return client.Database("test")
}

func insertUser(db *mongo.Database, user *User) {
	coll := db.Collection("user")
	coll.InsertOne(context.Background(), user)
}

func findUserByID(db *mongo.Database, userID string) (*User, error) {
	log.Printf("Try to find user: %s\n", userID)
	result := &User{}
	filter := bson.NewDocument(bson.EC.String("id", userID))
	err := db.Collection("user").FindOne(context.Background(), filter).Decode(result)
	return result, err
}

func main() {
	db := createDBConnection()
	user := User{ID: "xhocht", Name: "Thomas", Email: "thomas@hochbichler.at"}
	insertUser(db, &user)
	foundUser, err := findUserByID(db, "xhocht")

	if err != nil {
		log.Fatalf("Error finding user with id: %s, %v", "xhocht", err)
	}
	log.Printf("Found user: %#v", foundUser)
}
