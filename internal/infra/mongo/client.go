package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoConnection helps to handle the communication with mongoDB.
type MongoConnection struct {
	Collection *mongo.Collection
}

// NewMongoConnection initialization for the mongo collection abstraction.
func NewMongoConnection(ctx context.Context, dbName, collectionName string) (*MongoConnection, error) {
	db, err := newDB(ctx, dbName)
	if err != nil {
		return nil, err
	}

	return &MongoConnection{
		Collection: db.Collection(collectionName),
	}, nil
}

func (mc *MongoConnection) AddIndex(ctx context.Context, indexName string) (*MongoConnection, error) {
	index := mongo.IndexModel{
		Keys: bson.D{
			primitive.E{Key: indexName, Value: 1},
		},
	}

	if _, err := mc.Collection.Indexes().CreateOne(ctx, index); err != nil {
		return mc, err
	}

	return mc, nil
}

func newDB(ctx context.Context, dbName string) (*mongo.Database, error) {
	client, err := newMongoClient(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}

func newMongoClient(ctx context.Context) (*mongo.Client, error) {
	options := options.Client().
		SetReadPreference(readpref.Secondary()).
		ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

// Disconnect close the db connection safely to reduce problems like hanging db connexions.
func (mc *MongoConnection) Disconnect(ctx context.Context) error {
	return mc.Collection.Database().Client().Disconnect(ctx)
}
