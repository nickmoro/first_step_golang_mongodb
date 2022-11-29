package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	FieldStr  string `json:"fieldstr"`
	FieldInt  int    `json:"fieldint"`
	FieldBool bool   `json:"fieldbool"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("First_Database").Collection("First Collection")
	fmt.Println("Collection Type: ", reflect.TypeOf(col))

	oneDoc := MongoField{
		FieldStr:  "This is out first data",
		FieldInt:  3,
		FieldBool: true,
	}
	fmt.Println("oneDoc Type:", reflect.TypeOf(oneDoc))

	result, err := col.InsertOne(ctx, oneDoc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InsertOne() result type:", reflect.TypeOf(result))
	fmt.Println("InsertOne() result:", result)

	newID := result.InsertedID
	fmt.Println("InsertOne(), newID type:", reflect.TypeOf(newID))

}
