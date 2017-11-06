package models

//Credentials for login purposes
type Credentials struct {
	Id       int    `orm:"auto"`
	Username string `orm:"size(45)" ,json:"username"`
	Password string `orm:"size(45)" ,json:"password"`
}
