package mysql

import (
	"database/sql"
	"fmt"
	"sort"

	_ "github.com/go-sql-driver/mysql"
	"github.com/todolist/modellayer"
)

var server string = "localhost"
var port int = 3306
var databaseName string = "toDoListSchema"
var userName string = "devuser"
var password string = "abc123"

var dbConnection *sql.DB
var connected bool

type Mysql struct {}

func connect() {
	if !connected {
		fmt.Println("Creating new connection")
		var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, password, server, port, databaseName)
		var err error
		dbConnection, err = sql.Open("mysql", connectionString)

		if err != nil {
			panic(err.Error())
		}
		connected = true
	}
}

func (m Mysql) AddNewItem(item modellayer.Item) {
	fmt.Println("Adding new item, " + item.ToString())
	var insertString string = fmt.Sprintf("INSERT INTO Item VALUES ( '%d', '%s', '%d' )", item.ItemId, item.ItemName, item.GetDoneInt())
	performSimpleQuery(insertString)
}

func (m Mysql) UpdateItem(item modellayer.Item) {
	fmt.Println("Updating item, " + item.ToString())
	var updateString string = fmt.Sprintf("UPDATE Item SET itemName = '%s', done = '%d' WHERE itemId = '%d'", item.ItemName, item.GetDoneInt(), item.ItemId)
	performSimpleQuery(updateString)
}

func (m Mysql) DeleteItem(item modellayer.Item) {
	fmt.Println("Deleting item, " + item.ToString())
	var deleteString string = fmt.Sprintf("DELETE FROM Item WHERE itemId = '%d'", item.ItemId)
	performSimpleQuery(deleteString)
}

func (m Mysql) DeleteAllDoneItems() {
	fmt.Println("Deleting all done items")
	performSimpleQuery("DELETE FROM Item WHERE done = 1")
}

func (m Mysql) GetAllItems() []modellayer.Item {
	fmt.Println("Getting all items")

	connect()
	results, err := dbConnection.Query("SELECT itemId, itemName, done FROM Item")

	if err != nil {
		panic(err.Error())
	}

	var listOfItems []modellayer.Item
	for results.Next() {
        var item modellayer.Item
        err = results.Scan(&item.ItemId, &item.ItemName, &item.Done)
        if err != nil {
            panic(err.Error())
        }
		listOfItems = append(listOfItems, item)
    }
	
	results.Close()
	disconnect()

	if len(listOfItems) > 1 {
		sort.Slice(listOfItems, func(i int, j int) bool {
			return listOfItems[i].ItemId < listOfItems[j].ItemId
		})
	}

	return listOfItems
}

func (m Mysql) GetSingleItem(id int) modellayer.Item {
	fmt.Println("Getting a single item")

	connect()
	var item modellayer.Item
	err := dbConnection.QueryRow("SELECT itemId, itemName, done FROM Item WHERE itemId = ?", id).Scan(&item.ItemId, &item.ItemName, &item.Done)

	if err != nil && err.Error() == "sql: no rows in result set" {
		return modellayer.Item{}
	}

	if err != nil {
		panic(err.Error())
	}

	disconnect()
	return item
}

func performSimpleQuery(queryString string) {
	connect()

	query, err := dbConnection.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	query.Close()
	disconnect()
}

func disconnect() {
	dbConnection.Close()
	connected = false
}
