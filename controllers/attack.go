package controllers

import (
	"Golang_RPG/models"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icza/session"
	// "github.com/astaxie/beego/logs"
)

type BattleAttackController struct {
	beego.Controller
}

func attack(c *ChatController) {
	sess := session.Get(c.Ctx.Request)
	player := sess.Attr("bot").(models.Bots)
	enemy := sess.Attr("enemy").(models.Enemies)

	enemyCurrentHealth, _ := sess.Attr("enemyCurrentHealth").(int)

	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Intn(500)
	if random <= enemy.Agility {
		random = 0
	} else {
		random = 1
	}

	enemyCurrentHealth = enemyCurrentHealth - (player.Attack-enemy.Defense/100)*random

	if enemyCurrentHealth <= 0 {
		Win(c)
	} else {
		sess.SetAttr("enemyCurrentHealth", enemyCurrentHealth)
		EnemyTurn(c, enemy, player)
	}
}

func ChatAttack(c *ChatController) {
	attack(c)
}
