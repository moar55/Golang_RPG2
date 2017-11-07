package controllers

import (
	"Golang_RPG/models"
	"strconv"
	"strings"

	"Golang_RPG/errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icza/session"
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
	Message string       `json: "message"`
	Bot     *models.Bots `json: "yourBot"`
}

type SuccessWOBot struct {
	Message  string `json: "message"`
	BotError string `json: "error"`
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
		c.Ctx.ResponseWriter.WriteHeader(401)
	} else {
		sess := session.Get(c.Ctx.Request)
		sess.SetAttr("inBattle", false)
		sess.SetAttr("bot", bot)
		c.Data["json"] = &SuccessWBot{Message: "Welcome " + name + " !", Bot: &bot}
	}
}

func ChatLogin(username string, password string, c *ChatController) {

	o := orm.NewOrm()
	user := models.Users{Username: username, Password: password}
	err := o.Read(&user, "Username", "Password")
	user.Password = ""

	if err == orm.ErrNoRows {
		c.Data["json"] = &errors.WrongCredentials.Message
		c.Ctx.ResponseWriter.WriteHeader(401)

	} else {
		sess := session.NewSession()
		// /c.SetSession("id", user.Id)
		sess.SetAttr("id", user.Id)
		getBot(c, user.Id, user.Name, o)
		session.Add(sess, c.Ctx.ResponseWriter)
	}
	c.ServeJSON()
}
