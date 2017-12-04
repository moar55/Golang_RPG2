package controllers

import (
	"Golang_RPG/models"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type Briefing struct {
	MaxHealth      int
	CurrentHealth  int
	EnemyMaxHealth int
	EnemyHealth    int
}

func EnemyTurn(c *ChatController, enemy *models.Enemies, player *models.Bots) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")

	//TODO: Add boss/skill logic
	playerCurrentHealth, _ := session.Values["playerCurrentHealth"].(int)
	enemyCurrentHealth, _ := session.Values["enemyCurrentHealth"].(int)

	formula := enemy.Power

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	if playerCurrentHealth <= 0 {
		Lose(c)
	} else {
		session.Values["playerCurrentHealth"] = playerCurrentHealth
		session.Values["enemyCurrentHealth"] = enemyCurrentHealth
		c.Data["json"] = &Message{
			Message: player.Name + ": " + strconv.Itoa(playerCurrentHealth) + " / " + strconv.Itoa(player.Maxhp) +
				"    " + enemy.Name + ": " + strconv.Itoa(enemyCurrentHealth) + " / " + strconv.Itoa(enemy.Maxhp),
			Mode: "Turn",
		}
		session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
		c.ServeJSON()
	}
}

func DEnemyTurn(c *ChatController, enemy *models.Enemies, player *models.Bots) {
	//TODO: Add boss/skill logic
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")

	playerCurrentHealth, _ := session.Values["playerCurrentHealth"].(int)
	enemyCurrentHealth, _ := session.Values["enemyCurrentHealth"].(int)

	formula := enemy.Power / 2

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	fmt.Println("Your health ", playerCurrentHealth)
	if playerCurrentHealth <= 0 {
		DLose(c)
	} else {
		session.Values["playerCurrentHealth"] = playerCurrentHealth
		session.Values["enemyCurrentHealth"] = enemyCurrentHealth
		c.Data["json"] = &Message{
			Message: player.Name + ": " + strconv.Itoa(playerCurrentHealth) + " / " + strconv.Itoa(player.Maxhp) +
				"    " + enemy.Name + ": " + strconv.Itoa(enemyCurrentHealth) + " / " + strconv.Itoa(enemy.Maxhp),
			Mode: "DTurn",
		}
		session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
		c.ServeJSON()
	}
}
