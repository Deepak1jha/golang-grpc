package connection

import "go.mongodb.org/mongo-driver/mongo"

type Connection struct {
	Mongo *mongo.Client
}

func NewConnection(mongo *mongo.Client) *Connection {
	return &Connection{
		Mongo: mongo,
	}
}
