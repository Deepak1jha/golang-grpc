package connection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoConnect(mongoUri string) *mongo.Client {

	clientOpts := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Println("Not able to connect to Mongo: ", err)
	}

	// ping check
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println("Not able to ping to Mongo: ", err)
	}

	return client
}
