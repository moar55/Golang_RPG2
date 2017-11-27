package main

import (
	_ "Golang_RPG/routers"
	_ "Golang_RPG/scripts"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var wow = "hello"

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
	// globalSessions, _ = session.NewManager("memory", `{"cookieName":"session2", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	// go globalSessions.GC()
	beego.Run()
}
