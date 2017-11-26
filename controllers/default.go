package controllers

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"

	"github.com/astaxie/beego"
)

//MainController controls the root API functions
type MainController struct {
	beego.Controller
}

//TODO: Add an extra 'options' component which controls
//			the message options that the user can choose from

//Welcome is a welcoming struct
type Welcome struct {
	ServerStatus bool     `json:"serverStatus"`
	message      string   `json:"message"`
	Options      []string `json:"options"`
}

type Welcome2 struct {
	Message string `json:"message"`
	UUIDVal string `json:"uuid"`
}

//Get gets
func (c *MainController) Get() {
	if c.GetSession("userId") == nil {
		options := []string{"Yes", "No"}
		fmt.Println(options)
		u1 := uuid.NewV4()
		x := Welcome2{"Welcome adventurer! Did you happen to visit this realm before? If you would like to login please type 'login _ _' and fill the 2 spaces with your username and password otherwise type 'register _ _ _ _' and fill them with username password name age. ", u1.String()}
		// userID := session.Get("UserID")

		// }

		c.Data["json"] = &x
		// // TODO:Hey jude don't forget to uncomment
		// // l := logs.GetLogger()
		// // l.Println(x.ServerStatus)

	} else {
		options := []string{"Continue"}
		l := logs.GetLogger()
		l.Println(c.GetSession("userId"))
		y := fmt.Sprintf("Welcome, %d", c.GetSession("userId").(int))
		x := Welcome{true, y, options}
		c.Data["json"] = &x
	}

	c.ServeJSON()
}
