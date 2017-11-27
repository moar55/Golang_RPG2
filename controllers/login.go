package controllers

import (
	"Golang_RPG/models"
	"fmt"
	"strconv"
	"strings"

	"Golang_RPG/errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	} else {
		session, _ := store.Get(c.Ctx.Request, "session")
		fmt.Println("the session is", session)
		session.Values["inBattle"] = false
		session.Values["bot"] = bot
		c.Data["json"] = &SuccessWBot{Message: "Welcome " + name + " !", Bot: &bot}
	}
}

func ChatLogin(username string, password string, c *ChatController) {

	fmt.Println("In da login")
	o := orm.NewOrm()
	user := models.Users{Username: username, Password: password}
	err := o.Read(&user, "Username", "Password")
	user.Password = ""

	if err == orm.ErrNoRows {
		c.Data["json"] = &errors.WrongCredentials.Message
		c.Ctx.ResponseWriter.WriteHeader(401)

	} else {
		session, err := store.Get(c.Ctx.Request, "session")
		session.Options.HttpOnly = true
		// mysqlstore.MySQLStore.Options.HttpOnly = true

		// 	Path:     "/",
		// 	MaxAge:   86400 * 7,
		// 	HttpOnly: true,
		// }

		if err != nil {
			c.Data["json"] = &Message{Message: err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(500)
		}
		fmt.Println("setting da cookie boy")
		session.Values["id"] = user.Id
		fmt.Println("the session is", session.Values["id"])
		getBot(c, user.Id, user.Name, o)
		session.Save(c.Ctx.Request, c.Ctx.ResponseWriter.ResponseWriter)
	}
	// r.JSON(c.Ctx.ResponseWriter, http.StatusOK, map[string]string{"hello": "json"})

	c.ServeJSON()
}
