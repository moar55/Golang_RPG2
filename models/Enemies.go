package models

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Enemies struct {
	Id        int    `orm:"auto"`
	Name      string `orm:"size(45)" ,json:"name"`
	Type      int    `orm:"size(45)" ,json:"type"`
	Location  string `orm:"size(45)" ,json:"location"`
	Attack    int    `orm:"size(45)" ,json:"attack"`
	Defense   int    `orm:"size(45)" ,json:"defense"`
	Pp        int    `orm:"size(45)" ,json:"pp"`
	Agility   int    `orm:"size(45)" ,json:"agility"`
	Maxhp     int    `orm:"size(45)" ,json:"maxhp"`
	Fakka     int    `orm:"size(45)" ,json:"fakka"`
	Drop_item int    `orm:"size(45)" ,json:"drop_item"`
	Power     int    `orm:"size(45)" ,json:"power"`
}

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func TurnToEnemy(enemy orm.Params) Enemies {
	ID, _ := StrToInt(enemy["id"].(string))
	Type, _ := StrToInt(enemy["type"].(string))
	Attack, _ := StrToInt(enemy["attack"].(string))
	Defense, _ := StrToInt(enemy["defense"].(string))
	Pp, _ := StrToInt(enemy["pp"].(string))
	Agility, _ := StrToInt(enemy["agility"].(string))
	Maxhp, _ := StrToInt(enemy["maxhp"].(string))
	Fakka, _ := StrToInt(enemy["fakka"].(string))
	Power, _ := StrToInt(enemy["power"].(string))
	dp := enemy["drop_item"]
	L := enemy["location"]
	DropItem := -1
	Location := ""
	if dp == "" {
		DP, _ := StrToInt(dp.(string))
		DropItem = DP
	}
	if L == nil {
		Location = ""
	} else {
		Location = enemy["location"].(string)
	}

	return Enemies{
		Id:        ID,
		Name:      enemy["name"].(string),
		Type:      Type,
		Location:  Location,
		Attack:    Attack,
		Defense:   Defense,
		Pp:        Pp,
		Agility:   Agility,
		Maxhp:     Maxhp,
		Fakka:     Fakka,
		Drop_item: DropItem,
		Power:     Power,
	}
}
