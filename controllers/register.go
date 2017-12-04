package controllers

import (
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type myError struct {
	Message error `json:"message"`
}

type RegisterController struct {
	beego.Controller
}

func ChatRegister(username string, password string, name string, age int, c *ChatController) {

	o := orm.NewOrm()
	pass := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	hashedPass := string(hashedPassword[:])
	// TODO: handle error lol
	x := models.Users{
		Username: username,
		Password: hashedPass,
		Name:     name,
		Age:      age,
	}
	_, err := o.Insert(&x)
	if err != nil {
		c.Data["json"] = &myError{Message: err}
	} else {
		c.Data["json"] = &Message{Message: "Congratulations, you just registered! Welcome, " + x.Name + ". \n" +
			"You can use 'scan' to scan for nearby enemies or items, or you can type 'location' and enter your coordinates to look for shops",
			Mode: "Register",
		}
		session, _ := store.Get(c.Ctx.Request, "session")
		session.Values["id"] = x.Id
		session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
	}
	c.ServeJSON()
}
