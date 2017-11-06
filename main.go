package main

import (
	_ "Golang_RPG/routers"
	_ "Golang_RPG/scripts"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	// https://github.com/astaxie/beego/issues/1037
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		if ctx.Input.Method() == "OPTIONS" {
			ctx.Output.Header("Access-Control-Allow-Origin", "*")
			ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
			ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			ctx.Abort(200, "Hello")
		}
	})

	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
