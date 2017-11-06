package scan

import (
	"Golang_RPG/models"

	"github.com/astaxie/beego/orm"
)

type Item struct {
	Message string       `json:"message"`
	Item    models.Items `json:"item"`
}

func FoundItem(message string, item models.Items, botId int) *Item {
	o := orm.NewOrm()
	newItem := models.Inventory{
		BotId:  botId,
		ItemId: item.Id,
	}
	o.Insert(&newItem)
	return &Item{Message: message + item.Name, Item: item}
}
