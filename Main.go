package main

import (
	"github.com/todolist/integrationlayer"
	"fmt"
)

func main() {
	integrationlayer.GetSingleItem(1)

	listOfItems := integrationlayer.GetAllItems();

	for i, s := range listOfItems {
		fmt.Println(i, s.ToString())
	}
}
