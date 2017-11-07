package controllers

import (
	"Golang_RPG/controllers/scan"
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icza/session"
	// "github.com/astaxie/beego/logs"
)

type ScanController struct {
	beego.Controller
}

func ChatScan(c *ChatController) {
	sess := session.Get(c.Ctx.Request)
	o := orm.NewOrm()
	if sess.CAttr("id") != nil && sess.CAttr("bot") != nil {
		if sess.CAttr("inBattle") == false {
			rand.Seed(time.Now().UTC().UnixNano())
			random := rand.Intn(100)
			fmt.Println("Random number is ", random)
			//50 % chance of finding nothing
			if random < 50 {
				c.Data["json"] = scan.Nothing
			} else {
				//40% for a BATTLE!
				if random < 90 {
					//Battle stuff
					sess.SetAttr("inBattle", true)
					if sess.CAttr("inLocation") == true {
						//Found a BOSS enemy
						var enemies []orm.Params
						_, err := o.Raw("SELECT * FROM enemies WHERE type = ? order by rand() limit 1", "2").Values(&enemies)
						if err != nil {
							c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
						} else {
							fmt.Println(enemies[0])
							enemy := models.TurnToEnemy(enemies[0])
							sess.SetAttr("enemyCurrentHealth", enemy.Maxhp)
							sess.SetAttr("playerCurrentHealth", sess.CAttr("bot").(models.Bots).Maxhp)
							sess.SetAttr("enemy", enemy)
							c.Data["json"] = scan.EnterBattle("found random BOSS enemy! ", "boss", enemy)
						}
					} else {
						//Found a normal enemy
						var enemies []orm.Params
						_, err := o.Raw("SELECT * FROM enemies WHERE type = ? order by rand() limit 1", "1").Values(&enemies)
						if err != nil {
							c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
						} else {
							fmt.Println(enemies[0])
							enemy := models.TurnToEnemy(enemies[0])
							player := sess.CAttr("bot").(models.Bots)
							sess.SetAttr("enemyCurrentHealth", enemy.Maxhp)
							sess.SetAttr("playerCurrentHealth", player.Maxhp)
							fmt.Println(player)
							sess.SetAttr("enemy", enemy)
							c.Data["json"] = scan.EnterBattle("found a random normal enemy! ", "normal", enemy)
						}
					}
				} else {
					//10% chance of finding an item

					var items []orm.Params
					_, err := o.Raw("SELECT * FROM items WHERE type = ? order by rand() limit 1", "1").Values(&items)
					if err != nil {
						c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
					} else {
						fmt.Println(items[0])
						item := models.TurnToItem(items[0])
						bot := sess.CAttr("bot").(models.Bots)
						c.Data["json"] = scan.FoundItem("Found an item! ", item, bot.Id)
					}
				}
			}

		} else {
			c.Data["json"] = &errors.ErrorMessage{Message: "You're already in a battle!"}
		}
	} else {
		c.Data["json"] = &errors.ErrorMessage{Message: "Please login and/or create a bot"}
	}
	c.ServeJSON()
}
