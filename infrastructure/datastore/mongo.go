package datastore

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(dataBaseDomain, dataBaseName string) *mongo.Client  {
    database := dataBaseDomain + "/" + dataBaseName

    client, err := mongo.NewClient(options.Client().ApplyURI(database))
    if err != nil {
        log.Print("Error db: ",err.Error())
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Print("Error db: ",err.Error())
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Print("Error db: ",err.Error())
    }

    log.Print("Connected to MongoDB")
    return client
}

func GetCollection(client *mongo.Client, dataBaseName, collectionName string) *mongo.Collection {
    collection := client.Database(dataBaseName).Collection(collectionName)
    return collection
}
