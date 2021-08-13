package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/todolist/modellayer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connected bool 
var client *mongo.Client
var ctx context.Context
var itemsCollection *mongo.Collection

func connect() {
	if !connected {
		var err error
		client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		// defer client.Disconnect(ctx)
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}

		itemsCollection = client.Database("ToDoListDB").Collection("items")
		connected = true
	}
}

func AddNewItem(item modellayer.Item) {
	fmt.Println("Adding new item, " + item.ToString())
	
	connect()
	result, err := itemsCollection.InsertOne(ctx, item)
    if err != nil {
		log.Fatal(err)
    }
	fmt.Println(result);

	disconnect()
}

func UpdateItem(item modellayer.Item) {
	fmt.Println("Updating item, " + item.ToString())
	
	connect()

}

// func DeleteItem(item modellayer.Item) {
// 	fmt.Println("Deleting item, " + item.ToString())

// }

// func DeleteAllDoneItems() {
// 	fmt.Println("Deleting all done items")

// }

func GetAllItems() []modellayer.Item {
	fmt.Println("Getting all items")

	connect()
	results, err := itemsCollection.Find(ctx, bson.D{})

    if err != nil { 
		panic(err) 
	}
    defer results.Close(ctx)

    var listOfItems []modellayer.Item
    if err = results.All(ctx, &listOfItems); err != nil {
          panic(err)
    
	}
	results.Close(ctx)
	disconnect()
	return listOfItems
}

func disconnect() {
	client.Disconnect(ctx)
	connected = false
}

func GetSingleItem(id int) modellayer.Item {
	fmt.Println("Getting a single item")

	connect()
	filter := bson.M{"itemId":id}
	var item modellayer.Item
	err := itemsCollection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		log.Fatal(err)
	}
		
	disconnect()
	return item
}