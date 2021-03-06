package controllers

import (
	"Golang_RPG/models"
	"strconv"
	"strings"

	"Golang_RPG/errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	// "github.com/astaxie/beego/logs"
)

func StrToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

type LoginController struct {
	beego.Controller
}

type SuccessWBot struct {
	Message string       `json:"message"`
	Bot     *models.Bots `json:"bot"`
}

type SuccessWOBot struct {
	Message  string `json:"message"`
	BotError string `json:"boterr"`
}

func getBot(c *ChatController, id int, name string, o orm.Ormer) {

	bot := models.Bots{User_id: id}
	err := o.Read(&bot, "User_id")
	if err != nil {
		_err := ""
		if err == orm.ErrNoRows {
			_err = "You don't have a bot!"
		} else {
			_err = err.Error()
		}
		c.Data["json"] = &SuccessWOBot{Message: "Welcome " + name + " !", BotError: _err}
	} else {
		session, _ := store.Get(c.Ctx.Request, "session")
		session.Values["inBattle"] = false
		session.Values["bot"] = bot
		c.Data["json"] = &SuccessWBot{Message: "Welcome " + name + " !", Bot: &bot}
	}
}

func ChatLogin(username string, password string, c *ChatController) {
	o := orm.NewOrm()
	user := models.Users{Username: username}
	err := o.Read(&user, "Username")

	if err == orm.ErrNoRows {
		c.Data["json"] = &errors.WrongCredentials.Message
		c.Ctx.ResponseWriter.WriteHeader(401)
	} else {
		hashPass := []byte(user.Password)
		pass := []byte(password)
		err = bcrypt.CompareHashAndPassword(hashPass, pass)
		if err != nil {
			c.Data["json"] = &errors.WrongCredentials.Message
			c.Ctx.ResponseWriter.WriteHeader(401)
		} else {
			session, err := store.Get(c.Ctx.Request, "session")
			session.Options.HttpOnly = true
			user.Password = ""
			if err != nil {
				c.Data["json"] = &Message{Message: err.Error(), Mode: "Error"}
				c.Ctx.ResponseWriter.WriteHeader(500)
			}
			session.Values["id"] = user.Id
			getBot(c, user.Id, user.Name, o)
			session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
		}
	}
	c.ServeJSON()
}
