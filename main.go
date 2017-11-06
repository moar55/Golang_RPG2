package main

import (
	_ "Golang_RPG/routers"
	_ "Golang_RPG/scripts"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
