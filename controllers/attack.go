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
	player := c.GetSession("bot").(models.Bots)
	enemy := c.GetSession("enemy").(models.Enemies)

	enemyCurrentHealth, _ := c.GetSession("enemyCurrentHealth").(int)

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
		c.SetSession("enemyCurrentHealth", enemyCurrentHealth)
		EnemyTurn(c, enemy, player)
	}
}

func ChatAttack(c *ChatController) {
	attack(c)
}
