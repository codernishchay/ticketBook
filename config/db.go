package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client
var DB *mongo.Database

func DBConnect() {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://nishi:2su3cR3MXunGRzAM@cluster0.zjfve.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DBClient = client
	DB = client.Database("ticketBook")
	fmt.Println("Connected to Database")
}
