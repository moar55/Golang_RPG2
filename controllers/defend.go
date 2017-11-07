package controllers

import (
	"Golang_RPG/models"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icza/session"
	// "github.com/astaxie/beego/logs"
)

type BattleDefendController struct {
	beego.Controller
}

func ChatDefend(c *ChatController) {
	sess := session.Get(c.Ctx.Request)
	enemy := sess.Attr("enemy").(models.Enemies)
	player := sess.Attr("bot").(models.Bots)
	DEnemyTurn(c, enemy, player)
}
