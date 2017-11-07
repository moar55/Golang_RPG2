package controllers

import (
	"Golang_RPG/models"
	"strconv"

	"github.com/icza/session"
)

type win struct {
	Message string
	Enemy   models.Enemies
	Fakka   int
}

func Win(c *ChatController) {
	sess := session.Get(c.Ctx.Request)
	enemy := sess.Attr("enemy").(models.Enemies)
	sess.SetAttr("inBattle", false)
	c.Data["json"] = &Message{Message: "You won! You gained " + strconv.Itoa(enemy.Fakka) + " Fakka!"}
	c.ServeJSON()
}
