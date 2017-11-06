package models

import "github.com/astaxie/beego/orm"

type Items struct {
	Id            int `orm:"auto"`
	RequiredLevel int `orm:"size(45)" ,json:"required_level"`
	Name          string
	Description   string
	Race          string
	Price         int `orm:"size(45)" ,json:"price"`
	Type          int `json:"type"`
}

func TurnToItem(enemy orm.Params) Items {
	ID, _ := StrToInt(enemy["id"].(string))
	RequiredLevel, _ := StrToInt(enemy["required_level"].(string))
	Name, _ := enemy["name"].(string)
	Description, _ := enemy["description"].(string)
	Race, _ := enemy["race"].(string)
	Price, _ := StrToInt(enemy["price"].(string))

	return Items{
		Id:            ID,
		RequiredLevel: RequiredLevel,
		Name:          Name,
		Description:   Description,
		Race:          Race,
		Price:         Price,
	}
}
