package mongodb

import (
	"context"
	"go-base-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func ConnectMongoDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the primary to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")
	MongoClient = client
}

// GetCollection returns a MongoDB collection from the specified database and collection name
func GetCollection(collectionName string) *mongo.Collection {
	dbName := config.AppConfig.GetString("persistence.mongodb.name") // Retrieve the dbName from the config
	return MongoClient.Database(dbName).Collection(collectionName)
}
