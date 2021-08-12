package weblayer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/todolist/integrationlayer"
    "github.com/todolist/modellayer"
)

var baseUrl string = "api/item/"
var getAllUrl string = baseUrl + "getall"
var createUrl string = baseUrl + "create"
var updateUrl string = baseUrl + "update"
var deleteUrl string = baseUrl + "delete"
var deleteAllDoneUrl string = baseUrl + "deletealldone"

func CreateApiEndpoints() {
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
    listOfItems := integrationlayer.GetAllItems();
    c.IndentedJSON(http.StatusOK, listOfItems)
}

func create(c *gin.Context) {
    allowControll(c)

    var newItem modellayer.Item
    if err := c.BindJSON(&newItem); err != nil {
        return
    }

    integrationlayer.AddNewItem(newItem);
    c.IndentedJSON(http.StatusOK, newItem)
}

func update(c *gin.Context) {
    allowControll(c)

    var updatedItem modellayer.Item
    if err := c.BindJSON(&updatedItem); err != nil {
        return
    }

    integrationlayer.UpdateItem(updatedItem);
    c.IndentedJSON(http.StatusOK, updatedItem)
}

func delete(c *gin.Context) {
    allowControll(c)

    var deletedItem modellayer.Item
    if err := c.BindJSON(&deletedItem); err != nil {
        return
    }

    integrationlayer.DeleteItem(deletedItem);
    c.IndentedJSON(http.StatusOK, deletedItem)
}

func deleteAllDone(c *gin.Context) {
    allowControll(c)
    integrationlayer.DeleteAllDoneItems();
    c.IndentedJSON(http.StatusOK, "OK")
}


func allowControll(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
}
