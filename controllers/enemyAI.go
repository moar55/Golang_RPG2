package controllers

import (
	"Golang_RPG/models"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/icza/session"
	// "github.com/astaxie/beego/logs"
)

type Briefing struct {
	MaxHealth      int
	CurrentHealth  int
	EnemyMaxHealth int
	EnemyHealth    int
}

func EnemyTurn(c *ChatController, enemy models.Enemies, player models.Bots) {
	sess := session.Get(c.Ctx.Request)
	//TODO: Add boss/skill logic
	playerCurrentHealth, _ := sess.Attr("playerCurrentHealth").(int)
	enemyCurrentHealth, _ := sess.Attr("enemyCurrentHealth").(int)

	formula := enemy.Power

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	fmt.Println("Your health ", playerCurrentHealth)
	if playerCurrentHealth <= 0 {
		Lose(c)
	} else {
		sess.SetAttr("playerCurrentHealth", playerCurrentHealth)
		sess.SetAttr("enemyCurrentHealth", enemyCurrentHealth)
		c.Data["json"] = &Message{
			Message: player.Name + ": " + strconv.Itoa(playerCurrentHealth) + " / " + strconv.Itoa(player.Maxhp) +
				"    " + enemy.Name + ": " + strconv.Itoa(enemyCurrentHealth) + " / " + strconv.Itoa(enemy.Maxhp),
		}
		c.ServeJSON()
	}
}

func DEnemyTurn(c *ChatController, enemy models.Enemies, player models.Bots) {
	sess := session.Get(c.Ctx.Request)

	//TODO: Add boss/skill logic
	playerCurrentHealth, _ := sess.Attr("playerCurrentHealth").(int)
	enemyCurrentHealth, _ := sess.Attr("enemyCurrentHealth").(int)

	formula := enemy.Power / 2

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	fmt.Println("Your health ", playerCurrentHealth)
	if playerCurrentHealth <= 0 {
		DLose(c)
	} else {
		sess.SetAttr("playerCurrentHealth", playerCurrentHealth)
		sess.SetAttr("enemyCurrentHealth", enemyCurrentHealth)
		c.Data["json"] = &Message{
			Message: player.Name + ": " + strconv.Itoa(playerCurrentHealth) + " / " + strconv.Itoa(player.Maxhp) +
				"    " + enemy.Name + ": " + strconv.Itoa(enemyCurrentHealth) + " / " + strconv.Itoa(enemy.Maxhp),
		}
		c.ServeJSON()
	}
}
