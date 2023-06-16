package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBConnection handles the connection to the MongoDB database.
type MongoDBConnection struct {
	db *mongo.Collection
}

// NewMongoDBConnection creates a new MongoDBConnection instance.
func NewMongoDBConnection(ctx context.Context, collection string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/mecnave")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database("mecnave").Collection(collection), nil
}
