package modellayer

import "fmt"

type Item struct {
	ItemId   int
	ItemName string
	Done     bool
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
