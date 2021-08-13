package mysql

import (
	"fmt"
	"testing"

	"github.com/todolist/modellayer"
)

var globalTestReference *testing.T;

func TestCRUDOperations(t *testing.T) {
	globalTestReference = t
	
	item0 := modellayer.Item{ItemId: 100, ItemName: "Test item", Done: false}
	item1 := modellayer.Item{ItemId: 100, ItemName: "Update item name", Done: false}
	item2 := modellayer.Item{ItemId: 100, ItemName: "Update item Name and done", Done: true}

	assertItemIsNull(GetSingleItem(100))

	AddNewItem(item0)
	assertItemsAreEqual(item0, GetSingleItem(item0.ItemId))

	UpdateItem(item1)
	assertItemsAreEqual(item1, GetSingleItem(item1.ItemId))

	UpdateItem(item2)
	assertItemsAreEqual(item2, GetSingleItem(item2.ItemId))

	DeleteItem(item2)
	assertItemIsNull(GetSingleItem(102))
}

func TestGetAllItems(t *testing.T) {
	globalTestReference = t

	// Cannot run this test if there already exists items in the database.
	assertAreEqual(len(GetAllItems()), 0)

	item0 := modellayer.Item{ItemId: 100, ItemName: "Item A", Done: false}
	item1 := modellayer.Item{ItemId: 101, ItemName: "Item B", Done: false}
	item2 := modellayer.Item{ItemId: 102, ItemName: "Item C", Done: true}

	AddNewItem(item0)
	AddNewItem(item1)
	AddNewItem(item2)

	assertItemsAreEqual(item0, GetSingleItem(100))
	assertItemsAreEqual(item1, GetSingleItem(101))
	assertItemsAreEqual(item2, GetSingleItem(102))

	listOfItems := GetAllItems()
	assertAreEqual(len(GetAllItems()), 3)

	for _, s := range listOfItems {
		fmt.Println(s.ToString())
	}

	assertItemsAreEqual(item0, listOfItems[0])
	assertItemsAreEqual(item1, listOfItems[1])
	assertItemsAreEqual(item2, listOfItems[2])

	DeleteItem(item0)
	DeleteItem(item1)
	DeleteItem(item2)

	assertItemIsNull(GetSingleItem(100))
	assertItemIsNull(GetSingleItem(101))
	assertItemIsNull(GetSingleItem(102))

	assertAreEqual(len(GetAllItems()), 0)
}

func TestGetAllItemsReturnsItemsInItemIdOrder(t *testing.T) {
	globalTestReference = t

	listOfItems := GetAllItems();
	// Cannot run this test if there already exists items in the database.
	assertAreEqual(len(listOfItems), 0)

	item2 := modellayer.Item{ItemId: 102, ItemName: "Item A", Done: false}
	item0 := modellayer.Item{ItemId: 100, ItemName: "Item B", Done: false}
	item1 := modellayer.Item{ItemId: 101, ItemName: "Item C", Done: true}

	AddNewItem(item2)
	AddNewItem(item0)
	AddNewItem(item1)

	listOfItems = GetAllItems()
	for _, s := range listOfItems {
		fmt.Println(s.ToString())
	}

	assertItemsAreEqual(item0, listOfItems[0])
	assertItemsAreEqual(item1, listOfItems[1])
	assertItemsAreEqual(item2, listOfItems[2])

	DeleteItem(item0)
	DeleteItem(item1)
	DeleteItem(item2)

	assertAreEqual(len(GetAllItems()), 0)
}

func TestDeleteAllDone(t *testing.T) {
	globalTestReference = t

	// Cannot run this test if there already exists items in the database.
	assertAreEqual(len(GetAllItems()), 0)

	item0 := modellayer.Item{ItemId: 100, ItemName: "Item A", Done: true}
	item1 := modellayer.Item{ItemId: 101, ItemName: "Item B", Done: false}
	item2 := modellayer.Item{ItemId: 102, ItemName: "Item C", Done: true}

	AddNewItem(item0);
	AddNewItem(item1);
	AddNewItem(item2);

	assertAreEqual(len(GetAllItems()), 3)

	DeleteAllDoneItems()

	assertAreEqual(len(GetAllItems()), 1)

	assertItemsAreEqual(item1, GetSingleItem(item1.ItemId))

	DeleteItem(item1)
	assertAreEqual(len(GetAllItems()), 0)
}

func assertItemsAreEqual(itemA modellayer.Item, itemB modellayer.Item) {
	if !itemA.Equals(itemB) {
		globalTestReference.Fatalf("Items are unequal. " + itemA.ToString() + " | " +  itemB.ToString())
	}
}

func assertAreEqual(objectA int, objectB int) {
	if objectA != objectB {
		globalTestReference.Fatalf(fmt.Sprintf("Objects are not equal. %d != %d", objectA, objectB))
	}
}

func assertItemIsNull(itemA modellayer.Item) {
	if !itemA.IsNull() {
		globalTestReference.Fatalf("Item should no longer exist in the database | " + itemA.ToString())
	}
}