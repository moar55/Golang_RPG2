package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/gorilla/sessions"
)

type ChatController struct {
	beego.Controller
}

type Message struct {
	Message string `json:"message"`
	Mode    string `json:"mode"`
	Type    string `json:"type"`
}

// var store *redistore.RediStore

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
}

func loggedIn(session *sessions.Session) bool {
	return session.Values["id"] != nil
}

func (c *ChatController) Post() {
	if c.Ctx.Request.Header.Get("Authorization") != "" {
		decoder := json.NewDecoder(c.Ctx.Request.Body)
		var reqMessage Message
		decoder.Decode(&reqMessage)
		message := strings.Split(reqMessage.Message, " ")
		session, err := store.Get(c.Ctx.Output.Context.Request, "session")
		session.Options.HttpOnly = true

		if err != nil {
			c.Data["json"] = &Message{Message: err.Error(), Mode: "Error"}
			c.Ctx.ResponseWriter.WriteHeader(500)
		}

		fmt.Println(message)
		fmt.Println("the session is", session)

		switch message[0] {
		case "login":
			if loggedIn(session) {
				c.Ctx.ResponseWriter.WriteHeader(400)
				c.Data["json"] = &Message{Message: "You are already logged in, please logout first", Mode: "Error"}
				session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
				c.ServeJSON()
				return
			}
			ChatLogin(message[1], message[2], c)
		case "register":
			if loggedIn(session) {
				c.Data["json"] = &Message{Message: "You are already logged in", Mode: "Error"}
				c.Ctx.ResponseWriter.WriteHeader(400)
				session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
				c.ServeJSON()
				return
			}
			age, _ := strconv.Atoi(strings.Split(message[4], "\n")[0])
			ChatRegister(message[1], message[2], message[3], age, c)
		case "logout":
			if !loggedIn(session) {
				c.Data["json"] = &Message{Message: "You aren't logged ins", Mode: "Error"}
				c.Ctx.ResponseWriter.WriteHeader(400)
				session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
				c.ServeJSON()
			} else {
				session.Values["id"] = nil
				c.Data["json"] = &Message{Message: "Logged out", Mode: "Error"}
				session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
				c.ServeJSON()
			}
			return
		case "help":
			c.Data["json"] = &Message{Message: "register username password name age,\n login username password,\n scan to find enemies,\n bot <name> <race> to create bot,\n showShop to show nearest shop,\n buyItem <itemname> to buy an item, \n ", Mode: "Help"}
			c.ServeJSON()
			return
		case "search":
			if loggedIn(session) {
				break
			} else {
				c.Data["json"] = &Message{Message: "Locerr", Mode: "LocError"}
				c.ServeJSON()
			}
		default:
			if !loggedIn(session) {
				c.Data["json"] = &Message{Message: "Please either login or register, use help to get the required comments", Mode: "Error"}
				c.ServeJSON()
				return
			}
		}

		if loggedIn(session) {

			switch message[0] {
			case "bot":
				ChatBot(message[1], message[2], c)
			case "items":
				ChatGetItems(c)
			case "scan":
				ChatScan(c)
			case "attack":
				if session.Values["inBattle"] == true {
					ChatAttack(c)
				} else {
					c.Data["json"] = &Message{Message: "You aren't in a battle!", Mode: "Error"}
					c.Ctx.ResponseWriter.WriteHeader(400)
					c.ServeJSON()
				}
			case "item":
				if session.Values["inBattle"] == true {
					m := strings.Split(reqMessage.Message, "'")
					ChatItem(c, m[1])
				} else {
					c.Data["json"] = &Message{Message: "You aren't in a battle!", Mode: "Error"}
					c.Ctx.ResponseWriter.WriteHeader(400)
					c.ServeJSON()
				}
			case "defend":
				if session.Values["inBattle"] == true {
					ChatDefend(c)
				} else {
					c.Data["json"] = &Message{Message: "You aren't in a battle!", Mode: "Error"}
					c.Ctx.ResponseWriter.WriteHeader(400)
					c.ServeJSON()
				}
			case "search":
				if session.Values["inBattle"] == true {
					c.Data["json"] = &Message{Message: "You can't search in a battle!", Mode: "Error"}
					c.Ctx.ResponseWriter.WriteHeader(400)
					c.ServeJSON()
					return
				}
				lat, _ := strconv.ParseFloat(strings.Split(message[1], "\n")[0], 64)
				longt, _ := strconv.ParseFloat(strings.Split(message[2], "\n")[0], 64)
				fmt.Println("lat: ", lat, " long: ", longt)
				ChatSearch(lat, longt, c)
			case "showShop":
				ChatShop(c)
			case "buyItem":
				name := strings.Split(message[1], "'")[0]
				ChatBuy(c, name)
			default:
				c.Data["json"] = &Message{Message: "Incorrect input. Use help to get possible commands", Mode: "Error"}
				c.ServeJSON()
			}
		}
	} else {
		c.Data["json"] = &Message{Message: "Forbidden action: You have no uuid", Mode: "Error"}
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.ServeJSON()
	}
}
