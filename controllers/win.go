package controllers

import (
	"Golang_RPG/models"
	"strconv"
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
	c.Data["json"] = &Message{Message: "You won! You gained " + strconv.Itoa(enemy.Fakka) + " Fakka!"}
	c.ServeJSON()
}
