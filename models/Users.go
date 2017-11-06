package models

type Users struct {
	Id       int    `orm:"auto"`
	Username string `orm:"size(45)" ,json:"username"`
	Password string `orm:"size(45)" ,json:"password"`
	Name     string `orm:"size(45)" ,json:"name"`
	Age      int    `json:"age"`
}
