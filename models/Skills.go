package models

type Skills struct {
	Id             int    `orm:"auto"`
	Name           string `orm:"size(45)" ,json:"name"`
	Race           string `orm:"size(45)" ,json:"race"`
	Required_level int    `orm:"size(45)" ,json:"required_level"`
}
