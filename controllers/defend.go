package controllers

import (
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type BattleDefendController struct {
	beego.Controller
}

func ChatDefend(c *ChatController) {
	enemy := c.GetSession("enemy").(models.Enemies)
	player := c.GetSession("bot").(models.Bots)
	DEnemyTurn(c, enemy, player)
}
