package controllers

import (
	"Golang_RPG/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/logs"
)

type BattleItemController struct {
	beego.Controller
}
type duo struct {
	Name        string
	Description string
	Quantity    int
}

func ChatGetItems(c *ChatController) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	o := orm.NewOrm()
	bot, _ := session.Values["bot"].(*models.Bots)
	var items []orm.Params
	o.Raw("SELECT * FROM inventory WHERE bot_id = ?", bot.Id).Values(&items)
	var i = 0
	var str = "Your items : \n"
	var m = make(map[int]duo)
	for i < len(items) {
		var its []orm.Params
		x := items[i]["item_id"].(string)
		y, _ := strconv.Atoi(x)
		o.Raw("SELECT * FROM items WHERE id = ?", y).Values(&its)
		z, _ := strconv.Atoi(items[i]["item_id"].(string))
		m[z] = duo{Name: its[0]["name"].(string),
			Description: its[0]["description"].(string),
			Quantity:    m[z].Quantity + 1}
		i++
	}
	for _, value := range m {
		str += value.Name + " " + "  Description: " + value.Description + "  x" + strconv.Itoa(value.Quantity) + "\n"
	}
	c.Data["json"] = &Message{Message: str, Mode: "Items"}
	c.ServeJSON()
}

func ChatItem(c *ChatController, itemName string) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	o := orm.NewOrm()
	item := models.Items{Name: itemName}
	err := o.Read(&item, "Name")
	if err != nil {
		c.Data["json"] = &Message{Message: "That item does not exist!", Type: "Error"}
	} else {

		currentHealth, _ := session.Values["playerCurrentHealth"].(int)
		enemyHealth, _ := session.Values["enemyCurrentHealth"].(int)
		bot, _ := session.Values["bot"].(*models.Bots)
		enemy, _ := session.Values["enemy"].(*models.Enemies)

		inventory := &models.Inventory{ItemId: item.Id, BotId: bot.Id}
		err := o.Read(inventory, "ItemId", "BotId")
		if err != nil {
			c.Data["json"] = &Message{Message: "You do not have this item!", Type: "Error"}
		} else {

			currentHealth, enemyHealth, bot, enemy = useItem(itemName, currentHealth, enemyHealth, bot, enemy)
			session.Values["playerCurrentHealth"] = currentHealth
			session.Values["enemyCurrentHealth"] = enemyHealth
			session.Values["bot"] = bot
			session.Values["enemy"] = enemy

			var inventories []orm.Params
			o.Raw("SELECT * FROM inventory WHERE bot_id = ? AND item_id = ? limit 1", bot.Id, item.Id).Values(&inventories)
			x, _ := strconv.Atoi(inventories[0]["id"].(string))
			inv := &models.Inventory{Id: x}
			if _, err := o.Delete(inv); err != nil {
				c.Data["json"] = &Message{Message: err.Error(), Type: "Error"}
			} else {
				c.Data["json"] =
					&Message{Message: "You used " + itemName +
						"!\n " + item.Description +
						"\n  " + bot.Name + " : " + strconv.Itoa(currentHealth) + "/" + strconv.Itoa(bot.Maxhp) +
						"\n  " + enemy.Name + " : " + strconv.Itoa(enemyHealth) + " / " + strconv.Itoa(enemy.Maxhp), Type: "ItemUse"}
			}
		}
	}
	c.ServeJSON()

}

func useItem(itemName string, playerHealth int, enemyHealth int, bot *models.Bots, enemy *models.Enemies) (newCurrentHealth int, newEnemyHealth int, newBot *models.Bots, newEnemy *models.Enemies) {
	switch itemName {
	case ("Kofta Sandwich"):
		playerHealth += 200
		if playerHealth > bot.Maxhp {
			playerHealth = bot.Maxhp
		}
		break
	case ("Panini Sandwich"):
		playerHealth += int(float64(bot.Maxhp) * 0.3)
		if playerHealth > bot.Maxhp {
			playerHealth = bot.Maxhp
		}
		break
	case ("Chicken Ranch Sandwich"):
		playerHealth += int(float64(bot.Maxhp) * 0.5)
		if playerHealth > bot.Maxhp {
			playerHealth = bot.Maxhp
		}
		break
	case ("Cupcake"):
		break
	case ("Super Salad"):
		break
	case ("Health Potion"):
		playerHealth += int(float64(bot.Maxhp) * 0.15)
		if playerHealth > bot.Maxhp {
			playerHealth = bot.Maxhp
		}
	}
	return playerHealth, enemyHealth, bot, enemy
}
