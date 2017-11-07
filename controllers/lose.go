package controllers

import "github.com/icza/session"

type lose struct {
	Message string
}

func Lose(c *ChatController) {
	c.Data["json"] = &Message{Message: "You lose..."}
	c.ServeJSON()

}

func DLose(c *ChatController) {
	sess := session.Get(c.Ctx.Request)
	c.Data["json"] = &Message{Message: "You lose..."}
	sess.SetAttr("inBattle", false)
	c.ServeJSON()

}
