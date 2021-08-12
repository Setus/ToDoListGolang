package integrationlayer

import (
	"testing"

	"github.com/todolist/modellayer"
)


func TestCRUDOperations(t *testing.T) {
	
	var item0 modellayer.Item = modellayer.Item{ItemId: 100, ItemName: "Test item", Done: false}
	var item1 modellayer.Item = modellayer.Item{ItemId: 100, ItemName: "Update item name", Done: false}
	var item2 modellayer.Item = modellayer.Item{ItemId: 100, ItemName: "Update item Name and done", Done: true}


	if !GetSingleItem(100).IsNull() {
		t.Fatalf("Item should not yet exist in the database")
	}

	AddNewItem(item0)

	if !item0.Equals(GetSingleItem(100)) {
		t.Fatalf("Items are unequal")
	}

	UpdateItem(item1)

	if !item1.Equals(GetSingleItem(100)) {
		t.Fatalf("Items are unequal")
	}

	UpdateItem(item2)

	if !item2.Equals(GetSingleItem(100)) {
		t.Fatalf("Items are unequal")
	}

	DeleteItem(item2)

	if !GetSingleItem(100).IsNull() {
		t.Fatalf("Item should no longer exist in the database")
	}
}
