package controllers

import (
	"Golang_RPG/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type win struct {
	Message string
	Enemy   models.Enemies
	Fakka   int
}

func Win(c *ChatController) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	enemy := session.Values["enemy"].(*models.Enemies)
	session.Values["inBattle"] = false
	o := orm.NewOrm()
	bot := session.Values["bot"].(*models.Bots)
	bot.Fakka = bot.Fakka + enemy.Fakka
	fmt.Println(bot)
	o.Update(bot, "Fakka")
	c.Data["json"] = &Message{Message: "You won! You gained " + strconv.Itoa(enemy.Fakka) + " Fakka!", Mode: "Win"}
	session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
	c.ServeJSON()
}
