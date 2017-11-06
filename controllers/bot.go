package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BotController struct {
	beego.Controller
}
type Ayhabal struct {
	Id int `json:"id"`
}

func ChatBot(name string, race string, c *ChatController) {

	o := orm.NewOrm()
	_id := c.GetSession("id")
	if _id != nil {
		id := _id.(int)
		_bot := models.Bots{User_id: id}
		err := o.Read(&_bot)
		if err != nil {
			bot := models.Bots{
				Name:    name,
				Race:    race,
				Level:   1,
				User_id: id,
				Attack:  150,
				Defense: 10,
				Fakka:   100,
				Maxhp:   1000,
				Maxmp:   1000,
			}
			_, err := o.Insert(&bot)
			if err != nil {
				c.Data["json"] = &errors.Err{Message: err}
				c.Ctx.ResponseWriter.WriteHeader(401)
			} else {
				c.Data["json"] = &Message{Message: "Congratulations! You just created your bot, " + bot.Name}
			}
		} else {
			c.Data["json"] = &errors.HaveBot.Message
			c.Ctx.ResponseWriter.WriteHeader(401)
		}
	} else {
		c.Data["json"] = &errors.NotLoggedIn.Message
		c.Ctx.ResponseWriter.WriteHeader(401)
	}
	c.ServeJSON()
}
