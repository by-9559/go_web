package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Singleton struct {
	db *mongo.Client
}

//异步懒汉模式
// var once sync.Once 
var db  *mongo.Client

func useDatabase () *mongo.Client{
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://by9559:2FjNWHfZNnoHcoIY@serverdb.tgnra.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("link database ok!")
	return client
}



func GetDB() *mongo.Client {
	// once.Do(func() {
	// 	db = &Singleton{
	// 		db:useDatabase(),
	// 	}
	// })
	if db ==nil {
		db = useDatabase()
	}
	return db
}