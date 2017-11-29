package controllers

import (
	"Golang_RPG/models"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type BattleAttackController struct {
	beego.Controller
}

func attack(c *ChatController) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	player := session.Values["bot"].(*models.Bots)
	enemy := session.Values["enemy"].(*models.Enemies)
	enemyCurrentHealth, _ := session.Values["enemyCurrentHealth"].(int)

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
		session.Values["enemyCurrentHealth"] = enemyCurrentHealth
		EnemyTurn(c, enemy, player)
	}
}

func ChatAttack(c *ChatController) {
	attack(c)
}
