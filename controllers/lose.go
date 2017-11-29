package controllers

type lose struct {
	Message string
}

func Lose(c *ChatController) {
	c.Data["json"] = &Message{Message: "You lose..."}
	c.ServeJSON()

}

func DLose(c *ChatController) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	c.Data["json"] = &Message{Message: "You lose..."}
	session.Values["inBattle"] = false
	c.ServeJSON()

}
