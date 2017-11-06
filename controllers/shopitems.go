package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type NearestShopItemsController struct {
	beego.Controller
}

type Response2 struct {
	Response []*models.ShopItems `json:"message"`
}

func ChatShop(c *ChatController) {
	nearestShop := c.GetSession("nearShop")
	if nearestShop == nil {
		c.Data["json"] = &errors.SearchForShop.Message
		c.Ctx.ResponseWriter.WriteHeader(errors.SearchForShop.HTTPStatus)
	} else {
		o := orm.NewOrm()
		var shopItems []*models.ShopItems
		_, err := o.QueryTable("shop_items").Filter("LocationID", nearestShop).RelatedSel().All(&shopItems)
		if err != nil {
			c.Data["json"] = &Response{Message: err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(500)
		} else {
			var stringItems []string
			for _, x := range shopItems {
				stringItems = append(stringItems, fmt.Sprintf(
					"Name: %s, Required Level: %d, Description: %s, Race: %s, Price: %d\n",
					x.Item.Name, x.Item.RequiredLevel, x.Item.Description, x.Item.Race, x.Price))
			}
			c.Data["json"] = &Response{Message: "Available Items: " + strings.Join(stringItems, "")}
			// use the own below me for easier readabiliy (pure json :3)
			// c.Data["json"] = &Response2{Response: shopItems}
		}
		c.ServeJSON(true)
	}
}
