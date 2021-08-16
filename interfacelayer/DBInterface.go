package interfacelayer

import "github.com/todolist/modellayer"

type DBIntegrationInterface interface {
	AddNewItem(modellayer.Item)
	
	UpdateItem(modellayer.Item)
	
	GetSingleItem(int) modellayer.Item

	GetAllItems() []modellayer.Item

	DeleteItem(modellayer.Item)

	DeleteAllDoneItems()
}