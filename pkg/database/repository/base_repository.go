package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository[T any] struct {
	Client *mongo.Client
}

func NewBaseRepository[T any](client *mongo.Client) *BaseRepository[T] {
	return &BaseRepository[T]{
		Client: client,
	}
}

func (b *BaseRepository[T]) InsertOne(ctx context.Context, databaseName, collectionName string, entity interface{}) (primitive.ObjectID, error) {
	res, err := b.Client.Database(databaseName).Collection(collectionName).InsertOne(ctx, entity)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
