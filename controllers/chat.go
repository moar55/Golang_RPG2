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
}

// var store *redistore.RediStore

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	// storetemp, err := redistore.NewRediStore(10, "tcp", "localhost:8000", "", []byte("secret-key"))
	// // store, err := NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	//
	// if err != nil {
	// 	panic(err)
	// }
	// store, _ = mysqlstore.NewMySQLStore("b74fb0aa2d159d:847dba6e@tcp(us-cdbr-iron-east-05.cleardb.net:3306)/heroku_c82a81d5007f2fa?parseTime=true&loc=Local", "sessions", "/", 3600, []byte("supersecretkey"))
	// store2, _ := mysqlstore.NewMySQLStoreFromConnection(db, "sessions", "/", 86400*7)
	// fmt.Println(store2)
}

func (c *ChatController) Post() {
	// if c.Ctx.Request.Header.Get("Authorization") != "" {
	// }
	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var reqMessage Message
	decoder.Decode(&reqMessage)
	message := strings.Split(reqMessage.Message, " ")
	session, err := store.Get(c.Ctx.Output.Context.Request, "session")
	session.Options.HttpOnly = true

	if err != nil {
		c.Data["json"] = &Message{Message: err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(500)
	}

	fmt.Println(message)
	if session.Values["id"] == nil {
		fmt.Println("the session is", session)

		switch message[0] {
		case "login":
			fmt.Println("awesome")
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
	// } else {
	// 	c.Data["json"] = &Message{Message: "Forbidden action: You have no uuid"}
	// 	c.Ctx.ResponseWriter.WriteHeader(403)
	// 	c.ServeJSON()
	// }
}
