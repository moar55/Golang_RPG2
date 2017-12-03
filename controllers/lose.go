package controllers

type lose struct {
	Message string
}

func Lose(c *ChatController) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	c.Data["json"] = &Message{Message: "You lose...", Mode: "Lose"}
	session.Values["inBattle"] = false
	session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
	c.ServeJSON()

}

func DLose(c *ChatController) {
	session, _ := store.Get(c.Ctx.Output.Context.Request, "session")
	c.Data["json"] = &Message{Message: "You lose...", Mode: "Lose"}
	session.Values["inBattle"] = false
	session.Save(c.Ctx.Request, c.Ctx.ResponseWriter)
	c.ServeJSON()

}
