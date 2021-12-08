package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoContext struct {
}

func InitializeMongoDB() *MongoContext{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))  

    return &MongoContext{
        Client: client,
        MainDatabase: client.Database("")
    }
}
