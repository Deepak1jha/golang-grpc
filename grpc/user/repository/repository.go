package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"user/connection"
)

type Repository struct {
	userRef    *mongo.Collection
	connection *connection.Connection
}

func NewRepository(connection *connection.Connection) *Repository {
	return &Repository{
		connection: connection,
		userRef:    connection.Mongo.Database("DEV").Collection("user"),
	}
}
