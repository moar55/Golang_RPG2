package controllers

type lose struct {
	Message string
}

func Lose(c *ChatController) {
	c.Data["json"] = &Message{Message: "You lose..."}
	c.ServeJSON()

}

func DLose(c *ChatController) {
	c.Data["json"] = &Message{Message: "You lose..."}
	c.SetSession("inBattle", false)
	c.ServeJSON()

}
