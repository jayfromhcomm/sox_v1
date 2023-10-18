package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client is the MongoDB client that can be used across the application
var Client *mongo.Client

// Conversation represents the structure of your conversation
type Conversation struct {
	ID        string    `bson:"_id,omitempty"`
	UserInput string    `bson:"userInput"`
	BotOutput string    `bson:"botOutput"`
	Timestamp time.Time `bson:"timestamp"`
}

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read MongoDB URI from .env
	mongoURI := os.Getenv("MONGODB_URI")

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

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
	Client = client
}

// InsertConversation inserts a new conversation into the database
func InsertConversation(conversation Conversation) {
	collection := Client.Database("StarCommandDB").Collection("conversations")
	_, err := collection.InsertOne(context.TODO(), conversation)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllConversations fetches all conversations from the database
func GetAllConversations() []Conversation {
	var conversations []Conversation

	collection := Client.Database("StarCommandDB").Collection("conversations")

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var conv Conversation
		err := cur.Decode(&conv)
		if err != nil {
			log.Fatal(err)
		}

		conversations = append(conversations, conv)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	return conversations
}
