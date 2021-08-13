package main

import (
	// "github.com/todolist/weblayer"
	"fmt"

	// "github.com/todolist/modellayer"
	"github.com/todolist/mongodb"
)

func main() {
	// weblayer.CreateApiEndpoints()

	// listOfItems := mongodb.GetAllItems()

	// for _, s := range listOfItems {
	// 	fmt.Println(s.ToString())
	// }

	// item0 := modellayer.Item{ItemId: 101, ItemName: "Item C", Done: true}

	// mongodb.AddNewItem(item0)

	// listOfItems = mongodb.GetAllItems()

	// for _, s := range listOfItems {
	// 	fmt.Println(s.ToString())
	// }

	fmt.Println(mongodb.GetSingleItem(102).ToString())

}
