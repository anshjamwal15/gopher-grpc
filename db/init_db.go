package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName   = "test-go"
	db       *mongo.Database
	mongoCtx context.Context
)

func InitMongo() {

	mongoCtx = context.Background()
	clientOpt := options.Client().ApplyURI("mongodb+srv://anshjamwal:dZs17JyIgaLJ2Kmw@cluster0.fniyb.mongodb.net/demo?retryWrites=true&w=majority")
	client, err := mongo.Connect(mongoCtx, clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(mongoCtx, nil)

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	db = client.Database(dbName)
}
