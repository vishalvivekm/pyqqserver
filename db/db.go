package db


import (
	"context"
	"fmt"
	"log"
	"sync"
  
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  )
  var (
	once sync.Once
	client *mongo.Client
  )
  func InitMongoDB(mongoURI string) *mongo.Client {
	once.Do(func() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
	 log.Fatal("error occured init mongo ",err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Error pinging MongoDB: %v\n", err)
		
	  }
	  fmt.Println("Pinged! successfully connected to Mongo!")
	})
	return client
  }