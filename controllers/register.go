package controllers

import (
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type myError struct {
	Message error `json: "message"`
}

type RegisterController struct {
	beego.Controller
}

func ChatRegister(username string, password string, name string, age int, c *ChatController) {

	o := orm.NewOrm()
	x := models.Users{
		Username: username,
		Password: password,
		Name:     name,
		Age:      age,
	}
	_, err := o.Insert(&x)
	if err != nil {
		c.Data["json"] = &myError{Message: err}
	} else {
		c.Data["json"] = &Message{Message: "Congratulations, you just registered! Welcome, " + x.Name + ". \n" +
			"You can use 'scan' to scan for nearby enemies or items, or you can type 'location' and enter your coordinates to look for shops",
		}
	}
	c.ServeJSON()
}
