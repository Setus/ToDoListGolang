package tests

import (
	"fmt"
	"testing"

	"github.com/todolist/modellayer"
	"github.com/todolist/servicelayer"
)

var globalTestReference *testing.T;

var mysql string = "mysql"
var mongodb string = "mongodb"

func TestCRUDOperations(t *testing.T) {
	CRUDOperations(mysql, t)
	CRUDOperations(mongodb, t)
}

func TestGetAllItems(t *testing.T) {
	GetAllItems(mysql, t)
	GetAllItems(mongodb, t)
}

func TestGetAllItemsReturnsItemsInItemIdOrder(t *testing.T) {
	GetAllItemsReturnsItemsInItemIdOrder(mysql, t)
	GetAllItemsReturnsItemsInItemIdOrder(mongodb, t)
}

func TestDeleteAllDone(t *testing.T) {
	DeleteAllDone(mysql, t)
	DeleteAllDone(mongodb, t)
}

func CRUDOperations(databaseType string, t *testing.T) {
	globalTestReference = t
	
	item0 := modellayer.Item{ItemId: 100, ItemName: "Test item", Done: false}
	item1 := modellayer.Item{ItemId: 100, ItemName: "Update item name", Done: false}
	item2 := modellayer.Item{ItemId: 100, ItemName: "Update item Name and done", Done: true}

	servicelayer.InstantiateDatabase(databaseType)

	assertItemIsNull(servicelayer.GetSingleItem(100))

	servicelayer.AddNewItem(item0)
	assertItemsAreEqual(item0, servicelayer.GetSingleItem(item0.ItemId))

	servicelayer.UpdateItem(item1)
	assertItemsAreEqual(item1, servicelayer.GetSingleItem(item1.ItemId))

	servicelayer.UpdateItem(item2)
	assertItemsAreEqual(item2, servicelayer.GetSingleItem(item2.ItemId))

	servicelayer.DeleteItem(item2)
	assertItemIsNull(servicelayer.GetSingleItem(102))
}

func GetAllItems(databaseType string, t *testing.T) {
	globalTestReference = t

	servicelayer.InstantiateDatabase(databaseType)

	// Cannot run this test if there already exists items in the database.
	assertAreEqual(len(servicelayer.GetAllItems()), 0)

	item0 := modellayer.Item{ItemId: 100, ItemName: "Item A", Done: false}
	item1 := modellayer.Item{ItemId: 101, ItemName: "Item B", Done: false}
	item2 := modellayer.Item{ItemId: 102, ItemName: "Item C", Done: true}

	servicelayer.AddNewItem(item0)
	servicelayer.AddNewItem(item1)
	servicelayer.AddNewItem(item2)

	assertItemsAreEqual(item0, servicelayer.GetSingleItem(100))
	assertItemsAreEqual(item1, servicelayer.GetSingleItem(101))
	assertItemsAreEqual(item2, servicelayer.GetSingleItem(102))

	listOfItems := servicelayer.GetAllItems()
	assertAreEqual(len(servicelayer.GetAllItems()), 3)

	for _, s := range listOfItems {
		fmt.Println(s.ToString())
	}

	assertItemsAreEqual(item0, listOfItems[0])
	assertItemsAreEqual(item1, listOfItems[1])
	assertItemsAreEqual(item2, listOfItems[2])

	servicelayer.DeleteItem(item0)
	servicelayer.DeleteItem(item1)
	servicelayer.DeleteItem(item2)

	assertItemIsNull(servicelayer.GetSingleItem(100))
	assertItemIsNull(servicelayer.GetSingleItem(101))
	assertItemIsNull(servicelayer.GetSingleItem(102))

	assertAreEqual(len(servicelayer.GetAllItems()), 0)
}

func GetAllItemsReturnsItemsInItemIdOrder(databaseType string, t *testing.T) {
	globalTestReference = t

	servicelayer.InstantiateDatabase(databaseType)

	listOfItems := servicelayer.GetAllItems();
	// Cannot run this test if there already exists items in the database.
	assertAreEqual(len(listOfItems), 0)

	item2 := modellayer.Item{ItemId: 102, ItemName: "Item A", Done: false}
	item0 := modellayer.Item{ItemId: 100, ItemName: "Item B", Done: false}
	item1 := modellayer.Item{ItemId: 101, ItemName: "Item C", Done: true}

	servicelayer.AddNewItem(item2)
	servicelayer.AddNewItem(item0)
	servicelayer.AddNewItem(item1)

	listOfItems = servicelayer.GetAllItems()
	for _, s := range listOfItems {
		fmt.Println(s.ToString())
	}

	assertItemsAreEqual(item0, listOfItems[0])
	assertItemsAreEqual(item1, listOfItems[1])
	assertItemsAreEqual(item2, listOfItems[2])

	servicelayer.DeleteItem(item0)
	servicelayer.DeleteItem(item1)
	servicelayer.DeleteItem(item2)

	assertAreEqual(len(servicelayer.GetAllItems()), 0)
}

func DeleteAllDone(databaseType string, t *testing.T) {
	globalTestReference = t

	servicelayer.InstantiateDatabase(databaseType)

	// Cannot run this test if there already exists items in the database.
	assertAreEqual(len(servicelayer.GetAllItems()), 0)

	item0 := modellayer.Item{ItemId: 100, ItemName: "Item A", Done: true}
	item1 := modellayer.Item{ItemId: 101, ItemName: "Item B", Done: false}
	item2 := modellayer.Item{ItemId: 102, ItemName: "Item C", Done: true}

	servicelayer.AddNewItem(item0);
	servicelayer.AddNewItem(item1);
	servicelayer.AddNewItem(item2);

	assertAreEqual(len(servicelayer.GetAllItems()), 3)

	servicelayer.DeleteAllDoneItems()

	assertAreEqual(len(servicelayer.GetAllItems()), 1)

	assertItemsAreEqual(item1, servicelayer.GetSingleItem(item1.ItemId))

	servicelayer.DeleteItem(item1)
	assertAreEqual(len(servicelayer.GetAllItems()), 0)
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