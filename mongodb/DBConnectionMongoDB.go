package mongodb

import (
	"context"
	"fmt"
	"time"
	"sort"

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

type Mongodb struct {}

func connect() {
	if !connected {
		var err error
		client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			panic(err.Error())
		}
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			panic(err.Error())
		}
		// defer client.Disconnect(ctx)
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			panic(err.Error())
		}

		itemsCollection = client.Database("ToDoListDB").Collection("items")
		connected = true
	}
}

func (m Mongodb) AddNewItem(item modellayer.Item) {
	fmt.Println("Adding new item, " + item.ToString())
	
	connect()
	_, err := itemsCollection.InsertOne(ctx, item)
    if err != nil {
		panic(err.Error())
    }
	// fmt.Println(result);
	disconnect()
}

func (m Mongodb) UpdateItem(item modellayer.Item) {
	fmt.Println("Updating item, " + item.ToString())
	
	connect()
	filter := bson.M{"itemId" : item.ItemId}
	update := bson.D{
		{"itemId", item.ItemId}, 
		{"itemName", item.ItemName},
		{"done", item.Done}}

	_, err := itemsCollection.ReplaceOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	disconnect()
}

func (m Mongodb) DeleteItem(item modellayer.Item) {
	fmt.Println("Deleting item, " + item.ToString())

	connect()
	filter := bson.M{"itemId" : item.ItemId}

	_, err := itemsCollection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}

	disconnect()
}

func (m Mongodb) DeleteAllDoneItems() {
	fmt.Println("Deleting all done items")

	connect()
	filter := bson.M{"done" : true}

	_, err := itemsCollection.DeleteMany(ctx, filter)
	if err != nil {
		panic(err)
	}

	disconnect()
}

func (m Mongodb) GetAllItems() []modellayer.Item {
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

	if len(listOfItems) > 1 {
		sort.Slice(listOfItems, func(i int, j int) bool {
			return listOfItems[i].ItemId < listOfItems[j].ItemId
		})
	}

	results.Close(ctx)
	disconnect()
	return listOfItems
}

func (m Mongodb) GetSingleItem(id int) modellayer.Item {
	fmt.Println("Getting a single item")

	connect()
	filter := bson.M{"itemId" : id}
	var item modellayer.Item
	err := itemsCollection.FindOne(ctx, filter).Decode(&item)
	
	if err != nil && err.Error() == "mongo: no documents in result" {
		return modellayer.Item{}
	}
	if err != nil {
		panic(err.Error())
	}
	disconnect()
	return item
}

func disconnect() {
	client.Disconnect(ctx)
	connected = false
}