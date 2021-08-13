package modellayer

import "fmt"

type Item struct {
	ItemId   int 	`json:"itemId" bson:"itemId"`
	ItemName string `json:"itemName" bson:"itemName"`
	Done     bool 	`json:"done" bson:"done"`
}

func (item Item) ToString() string {
	return fmt.Sprintf("itemId: %d, itemName: %s, done: %t", item.ItemId, item.ItemName, item.Done)
}

func (item Item) GetDoneInt() int {
	if item.Done {
		return 1
	} else {
		return 0
	}
}

func (item Item) Equals(itemToCompare Item) bool {
	if item.ItemId != itemToCompare.ItemId {
		return false
	}

	if item.ItemName != itemToCompare.ItemName {
		return false
	}

	if item.Done != itemToCompare.Done {
		return false
	}

	return true
}

func (item Item) IsNull() bool {
	return item.ItemName == ""
}
