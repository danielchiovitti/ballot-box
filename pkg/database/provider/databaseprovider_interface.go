package provider

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseProviderInterface interface {
	GetDb() (*mongo.Client, error)
}
