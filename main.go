package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func CheckErr(str string, err error) {
	if err != nil {
		log.Fatalln(str, err)
	}
}

// Prints all elements in mongo.Collection
func PrintlnAllElements(collection *mongo.Collection, ctx context.Context) error {
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			return err
		}
		log.Println(result)
	}
	if err := cursor.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	CheckErr("Server connection error:", err)
	CheckErr("Pinging the server error:", client.Ping(context.Background(), readpref.Primary()))
	defer func() {
		log.Println("Will disconnect from database")
		CheckErr("Server disconnecting error:", client.Disconnect(ctx))
		log.Println("Disconnected from database")
	}()
	collection := client.Database("First_Database").Collection("First Collection")
	log.Println("Connected to database")

	// Printing all elements
	PrintlnAllElements(collection, ctx)

	// // Deleting many by filter
	// filter := bson.D{{Key: "fieldbool", Value: true}}
	// result, err := collection.DeleteMany(ctx, filter)
	// CheckErr("DeleteMany error:", err)
	// log.Println("Deleted", result.DeletedCount, "elements")

	// // Inserting many
	// docsToInsert := []interface{}{
	// 	bson.D{
	// 		{Key: "fieldstr", Value: "first data"},
	// 		{Key: "fieldint", Value: 3},
	// 		{Key: "fieldbool", Value: true},
	// 	},
	// 	bson.D{
	// 		{Key: "fieldstr", Value: "second data"},
	// 		{Key: "fieldint", Value: 15},
	// 		{Key: "fieldbool", Value: false},
	// 	},
	// }
	// result, err := collection.InsertMany(ctx, docsToInsert)
	// CheckErr("InsertMany error:", err)
	// log.Println("Number of documents inserted:", len(result.InsertedIDs))
	// log.Println(result.InsertedIDs...)

}
