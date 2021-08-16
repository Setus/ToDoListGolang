package servicelayer

import (
	"github.com/todolist/interfacelayer"
	"github.com/todolist/modellayer"
	"github.com/todolist/mysql"
	"github.com/todolist/mongodb")

var dbConnection interfacelayer.DBIntegrationInterface

func InstantiateDatabase(databaseType string) {
	if databaseType == "mysql" {
		setDatabaseType(mysql.Mysql{})
	} else {
		setDatabaseType(mongodb.Mongodb{})
	}
}

func setDatabaseType(dbc interfacelayer.DBIntegrationInterface) {
	dbConnection = dbc
}

func AddNewItem(item modellayer.Item) {
	dbConnection.AddNewItem(item)
}

func UpdateItem(item modellayer.Item) {
	dbConnection.UpdateItem(item)
}

func GetSingleItem(itemId int) modellayer.Item {
	return dbConnection.GetSingleItem(itemId)
}

func GetAllItems() []modellayer.Item {
	return dbConnection.GetAllItems()
}

func DeleteItem(item modellayer.Item) {
	dbConnection.DeleteItem(item)
}

func DeleteAllDoneItems() {
	dbConnection.DeleteAllDoneItems()
}