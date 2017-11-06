package models

type ShopItems struct {
	Id         int `orm:"auto"`
	LocationId int
	Item       *Items `orm:"rel(one)"`
	Price      int
}
