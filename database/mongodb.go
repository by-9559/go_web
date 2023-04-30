package mongodb

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
    Client *mongo.Client
}

var db *mongo.Client

func useDatabase() *mongo.Client {
    serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
    clientOptions := options.Client().
        ApplyURI("mongodb+srv://by9559:2FjNWHfZNnoHcoIY@serverdb.tgnra.mongodb.net/mydb?retryWrites=true&w=majority").
        SetServerAPIOptions(serverAPIOptions)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("link database ok!")
    return client
}

func GetDB() *mongo.Client {
    if db == nil {
        db = useDatabase()
    }
    return db
}
