package config

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
	mongoURI            string
)

func loadEnvAndGetMongoURI() string {
	return os.Getenv("URI_MONGO")
}


func getMongoClient() (*mongo.Client, error){
	mongoOnce.Do(func() {

		mongoURI = loadEnvAndGetMongoURI()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(mongoURI)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil{
			clientInstanceError = err 
			return  
		}
		
		err = client.Ping(ctx, nil)
		if err != nil{
			clientInstanceError = err 
			return  
		}

		clientInstance = client
		log.Println("Connect to MongoDB successfully")
	})

	return clientInstance, clientInstanceError
	
}

func SetupDB() *mongo.Database{
	
	client, err := getMongoClient()

	if err != nil{
		log.Fatalf("Error connect mongo: %v", err)
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		log.Fatal("MONGO_DB_NAME not set in .env")
	}

	database := client.Database(dbName)

	log.Println("MongoDB database selected:", dbName)
	return database
}