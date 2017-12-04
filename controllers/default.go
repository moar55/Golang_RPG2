package controllers

import (
	"fmt"

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
	ServerStatus bool
	Message      string   `json:"message"`
	Options      []string `json:"options"`
}

//Welcome2 does stuff
type Welcome2 struct {
	Message string `json:"message"`
	UUIDVal string `json:"uuid"`
}

//Get gets
func (c *MainController) Get() {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")

	if session.Values["userId"] == nil {
		options := []string{"Yes", "No"}
		fmt.Println(options)
		u1 := uuid.NewV4()
		x := Welcome2{"Welcome adventurer! Did you happen to visit this realm before? If you would like to login please type 'login _ _' and fill the 2 spaces with your username and password otherwise type 'register _ _ _ _' and fill them with username password name age. ", u1.String()}
		c.Data["json"] = &x
	} else {
		options := []string{"Continue"}
		y := session.Values["userId"].(string)
		x := Welcome{true, y, options}
		c.Data["json"] = &x
	}

	c.ServeJSON()
}
