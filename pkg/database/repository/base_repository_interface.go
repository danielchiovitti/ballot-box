package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseRepositoryInterface[T any] interface {
	InsertOne(ctx context.Context, databaseName, collectionName string, entity interface{}) (primitive.ObjectID, error)
}
