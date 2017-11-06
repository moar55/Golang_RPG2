package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

type Message struct {
	Message string `json:"message"`
}

func (c *ChatController) Post() {
	if c.Ctx.Request.Header.Get("Authorization") != "" {

		message := strings.Split(c.GetString("message"), " ")

		if c.GetSession("id") == nil {
			switch message[0] {
			case "login":
				ChatLogin(message[1], message[2], c)
			case "register":
				age, _ := strconv.Atoi(strings.Split(message[4], "\n")[0])
				ChatRegister(message[1], message[2], message[3], age, c)
			case "help":
				c.Data["json"] = &Message{Message: "register username password name age, login username password"}
				c.ServeJSON()
			default:
				c.Data["json"] = &Message{Message: "Please either login or register, use help to get the required comments"}
				c.ServeJSON()
			}
		} else {

			switch message[0] {
			case "bot":
				ChatBot(message[1], message[2], c)
			case "scan":
				ChatScan(c)
			case "attack":
				ChatAttack(c)
			case "defend":
				ChatDefend(c)
			case "search":
				lat, _ := strconv.ParseFloat(strings.Split(message[1], "\n")[0], 64)
				longt, _ := strconv.ParseFloat(strings.Split(message[2], "\n")[0], 64)
				fmt.Println("lat: ", lat, " long: ", longt)
				ChatSearch(lat, longt, c)
			case "showShop":
				ChatShop(c)
			case "buyItem":
				name := strings.Split(c.GetString("message"), "'")[1]
				ChatBuy(c, name)
			default:
				c.Data["json"] = &Message{Message: "Incorrect input"}
				c.ServeJSON()
			}
		}
	} else {
		c.Data["json"] = &Message{Message: "Forbidden action: You have no uuid"}
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.ServeJSON()
	}
}
