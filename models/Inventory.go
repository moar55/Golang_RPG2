package models

type Inventory struct {
	Id     int `orm:"auto"`
	BotId  int `json:"bot_id"`
	ItemId int `json:"bot_id"`
}
