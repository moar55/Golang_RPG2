package models

type Locations struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(45)" ,json:"name"`
	Type string `orm:"size(45)" ,json:"type"`
}
