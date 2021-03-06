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
	// "github.com/astaxie/beego/logs"
)

type ScanController struct {
	beego.Controller
}

func ChatScan(c *ChatController) {
	o := orm.NewOrm()
	session, _ := store.Get(c.Ctx.Request, "session")
	fmt.Println("the session is", session)

	if session.Values["bot"] != nil {
		if session.Values["inBattle"] == false {
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
					session.Values["inBattle"] = true
					if session.Values["inLocation"] == true {
						//Found a BOSS enemy
						var enemies []orm.Params
						_, err := o.Raw("SELECT * FROM enemies WHERE type = ? order by rand() limit 1", "2").Values(&enemies)
						if err != nil {
							c.Data["json"] = &errors.ErrorMessage{Message: err.Error()}
						} else {
							fmt.Println(enemies[0])
							enemy := models.TurnToEnemy(enemies[0])
							session.Values["enemyCurrentHealth"] = enemy.Maxhp
							session.Values["playerCurrentHealth"] = session.Values["bot"].(*models.Bots).Maxhp
							session.Values["enemy"] = enemy
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
							player := session.Values["bot"].(*models.Bots)
							session.Values["enemyCurrentHealth"] = enemy.Maxhp
							session.Values["playerCurrentHealth"] = player.Maxhp
							fmt.Println(player)
							session.Values["enemy"] = enemy
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
						bot := session.Values["bot"].(*models.Bots)
						c.Data["json"] = scan.FoundItem("Found an item! ", item, bot.Id)
					}
				}
			}

		} else {
			c.Data["json"] = &errors.ErrorMessage{Message: "You're already in a battle!"}
		}
	} else {
		c.Data["json"] = &errors.ErrorMessage{Message: "Please create a bot"}
	}
	session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)

	c.ServeJSON()
}
