package weblayer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties"
	"github.com/todolist/modellayer"
	"github.com/todolist/servicelayer"
)

var baseUrl string = "api/item/"
var getAllUrl string = baseUrl + "getall"
var createUrl string = baseUrl + "create"
var updateUrl string = baseUrl + "update"
var deleteUrl string = baseUrl + "delete"
var deleteAllDoneUrl string = baseUrl + "deletealldone"

func CreateApiEndpoints() {
    servicelayer.InstantiateDatabase(readDatabaseSetting())
    
    router := gin.Default()

    router.GET(getAllUrl, getAll)

    router.OPTIONS(createUrl, allowControll)
    router.POST(createUrl, create)
    
    router.OPTIONS(updateUrl, allowControll)
    router.POST(updateUrl, update)
    
    router.OPTIONS(deleteUrl, allowControll)
    router.DELETE(deleteUrl, delete)

    router.OPTIONS(deleteAllDoneUrl, allowControll)
    router.DELETE(deleteAllDoneUrl, deleteAllDone)

    router.Run("localhost:21561")
}

func getAll(c *gin.Context) {
    allowControll(c)
    listOfItems := servicelayer.GetAllItems();
    c.IndentedJSON(http.StatusOK, listOfItems)
}

func create(c *gin.Context) {
    allowControll(c)

    var newItem modellayer.Item
    if err := c.BindJSON(&newItem); err != nil {
        return
    }

    servicelayer.AddNewItem(newItem);
    c.IndentedJSON(http.StatusOK, newItem)
}

func update(c *gin.Context) {
    allowControll(c)

    var updatedItem modellayer.Item
    if err := c.BindJSON(&updatedItem); err != nil {
        return
    }

    servicelayer.UpdateItem(updatedItem);
    c.IndentedJSON(http.StatusOK, updatedItem)
}

func delete(c *gin.Context) {
    allowControll(c)

    var deletedItem modellayer.Item
    if err := c.BindJSON(&deletedItem); err != nil {
        return
    }

    servicelayer.DeleteItem(deletedItem);
    c.IndentedJSON(http.StatusOK, deletedItem)
}

func deleteAllDone(c *gin.Context) {
    allowControll(c)
    servicelayer.DeleteAllDoneItems();
    c.IndentedJSON(http.StatusOK, "OK")
}

func allowControll(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
}

func readDatabaseSetting() string {
    p := properties.MustLoadFile("config.properties", properties.UTF8)
	databaseType := p.MustGetString("DatabaseType")
	fmt.Println("The databaseType is: " + databaseType)
    return databaseType
}