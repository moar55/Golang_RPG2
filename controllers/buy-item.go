package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BuyItemController struct {
	beego.Controller
}

func ChatBuy(c *ChatController, name string) {
	if c.GetSession("nearShop") == nil {
		fmt.Print(c.GetSession("nearShop"))
		c.Data["json"] = &Message{Message: "search again for shop"}
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.ServeJSON()
		return
	}
	var itemName string = name
	o := orm.NewOrm()
	item := models.Items{Name: itemName}
	err := o.Read(&item, "Name")
	if err == orm.ErrNoRows {
		c.Data["json"] = &Response{Message: "No item with such name"}
	} else {
		var shopItems []*models.ShopItems
		o.QueryTable("shop_items").Filter("LocationID", c.GetSession("nearShop")).RelatedSel().All(&shopItems)
		// could have used a join here, too late tho :/
		var shopItem *models.ShopItems
		for _, x := range shopItems {
			fmt.Println(x.Item.Name)
			if x.Item.Name == itemName {
				shopItem = x
				break
			}
		}
		if shopItem == nil {
			c.Data["json"] = &errors.ItemNotFound.Message
			c.Ctx.ResponseWriter.WriteHeader(errors.ItemNotFound.HTTPStatus)
		} else {
			bot := c.GetSession("bot").(models.Bots)
			if bot.Fakka <= shopItem.Price {
				c.Data["json"] = &errors.NoEnoughFakka.Message
				c.Ctx.ResponseWriter.WriteHeader(errors.NoEnoughFakka.HTTPStatus)
			} else {
				bot.Fakka = bot.Fakka - shopItem.Price
				fmt.Println("Bot fakka: ", bot.Fakka, " Item Price: ", shopItem.Price)
				_, err = o.Update(&bot, "Fakka")
				if err != nil {
					c.Data["json"] = &Response{Message: err.Error()}
					c.Ctx.ResponseWriter.WriteHeader(500)
				} else {
					c.Data["json"] = &Response{Message: "Done!"}
					c.SetSession("nearShop", nil)
				}
			}
		}
	}
	c.ServeJSON(true)

}
