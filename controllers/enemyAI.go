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

func EnemyTurn(c *ChatController, enemy models.Enemies, player models.Bots) {
	//TODO: Add boss/skill logic
	playerCurrentHealth, _ := c.GetSession("playerCurrentHealth").(int)
	enemyCurrentHealth, _ := c.GetSession("enemyCurrentHealth").(int)

	formula := enemy.Power

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	fmt.Println("Your health ", playerCurrentHealth)
	if playerCurrentHealth <= 0 {
		Lose(c)
	} else {
		c.SetSession("playerCurrentHealth", playerCurrentHealth)
		c.SetSession("enemyCurrentHealth", enemyCurrentHealth)
		c.Data["json"] = &Message{
			Message: player.Name + ": " + strconv.Itoa(playerCurrentHealth) + " / " + strconv.Itoa(player.Maxhp) +
				"    " + enemy.Name + ": " + strconv.Itoa(enemyCurrentHealth) + " / " + strconv.Itoa(enemy.Maxhp),
		}
		c.ServeJSON()
	}
}

func DEnemyTurn(c *ChatController, enemy models.Enemies, player models.Bots) {
	//TODO: Add boss/skill logic
	playerCurrentHealth, _ := c.GetSession("playerCurrentHealth").(int)
	enemyCurrentHealth, _ := c.GetSession("enemyCurrentHealth").(int)

	formula := enemy.Power / 2

	playerCurrentHealth =
		// playerCurrentHealth - (enemy.Power * (100 - (player.Defense / 500)))
		playerCurrentHealth - formula
	fmt.Println("Your health ", playerCurrentHealth)
	if playerCurrentHealth <= 0 {
		DLose(c)
	} else {
		c.SetSession("playerCurrentHealth", playerCurrentHealth)
		c.SetSession("enemyCurrentHealth", enemyCurrentHealth)
		c.Data["json"] = &Message{
			Message: player.Name + ": " + strconv.Itoa(playerCurrentHealth) + " / " + strconv.Itoa(player.Maxhp) +
				"    " + enemy.Name + ": " + strconv.Itoa(enemyCurrentHealth) + " / " + strconv.Itoa(enemy.Maxhp),
		}
		c.ServeJSON()
	}
}
