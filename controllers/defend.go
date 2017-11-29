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
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	enemy := session.Values["enemy"].(*models.Enemies)
	player := session.Values["bot"].(*models.Bots)
	DEnemyTurn(c, enemy, player)
}
