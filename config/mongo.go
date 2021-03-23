package config

import (
	"context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


func GetClient() (*mongo.Client){
    
    clientOptions := options.Client().ApplyURI("mongodb://host.docker.internal:9002")
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
	}

	return client
	
}