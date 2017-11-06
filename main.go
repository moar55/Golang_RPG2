package main

import (
	_ "Golang_RPG/routers"
	_ "Golang_RPG/scripts"

	_ "github.com/astaxie/beego/session/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	// https://github.com/astaxie/beego/issues/1037
	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		if ctx.Input.Method() == "OPTIONS" {
			ctx.Output.Header("Access-Control-Allow-Origin", "*")
			ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
			ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Abort(200, "Hello")
		}
	})
	beego.Run()
}
