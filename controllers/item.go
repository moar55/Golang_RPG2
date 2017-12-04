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
		c.Data["json"] = &Message{Message: "You don't have that item", Type: "Error"}
	} else {
		currentHealth, _ := session.Values["playerCurrentHealth"].(int)
		enemyHealth, _ := session.Values["enemyCurrentHealth"].(int)
		bot, _ := session.Values["bot"].(*models.Bots)
		enemy, _ := session.Values["enemy"].(*models.Enemies)

		currentHealth, enemyHealth, bot, enemy = useItem(itemName, currentHealth, enemyHealth, bot, enemy)
		session.Values["playerCurrentHealth"] = currentHealth
		session.Values["enemyCurrentHealth"] = enemyHealth
		session.Values["bot"] = bot
		session.Values["enemy"] = enemy
		c.Data["json"] =
			&Message{Message: "You used " + itemName +
				"!\n " + item.Description +
				"\n  " + bot.Name + " : " + strconv.Itoa(currentHealth) + "/" + strconv.Itoa(bot.Maxhp) +
				"\n  " + enemy.Name + " : " + strconv.Itoa(enemyHealth) + " / " + strconv.Itoa(enemy.Maxhp), Type: "ItemUse"}
	}
	c.ServeJSON()

}

func useItem(itemName string, playerHealth int, enemyHealth int, bot *models.Bots, enemy *models.Enemies) (newCurrentHealth int, newEnemyHealth int, newBot *models.Bots, newEnemy *models.Enemies) {
	switch itemName {
	case ("Kofta Sandwich"):
		break
	case ("Panini Sandwich"):
		break
	case ("Chicken Ranch Sandwich"):
		break
	case ("Cupcake"):
		break
	case ("Super Salad"):
		break
	}
	return playerHealth, enemyHealth, bot, enemy
}
