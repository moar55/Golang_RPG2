package models

//RegisterCredentials for the registration process
type RegisterCredentials struct {
	Username string `orm:"size(100)" ,json:"username"`
	Password string `orm:"size(100)" ,json:"password"`
	Name     string `orm:"size(100)" ,json:"name"`
	Age      int    `json:"age"`
}
