package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://vbhawsar:YOURPWD@cluster0.7nynfbg.mongodb.net/?retryWrites=true&w=majority"

const dbName = "Netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// CONNECT WITH MongoDB

func init(){
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//Connect to MongoDB
	client , err := mongo.Connect(context.TODO() , clientOption)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//Collection instance 
	fmt.Println("Collection instance is ready")
}

