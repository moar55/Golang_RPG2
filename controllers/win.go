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
	enemy := c.GetSession("enemy").(models.Enemies)
	c.SetSession("inBattle", false)
	c.Data["json"] = &Message{Message: "You won! You gained " + strconv.Itoa(enemy.Fakka) + " Fakka!"}
	c.ServeJSON()
}
