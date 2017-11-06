package routers

import (
	"Golang_RPG/controllers"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var checkForAuthorization = func(ctx *context.Context) {

	// ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	subUrl := strings.Split(ctx.Request.URL.Path, "/")[1]
	println(subUrl)
	if subUrl == "login" || subUrl == "register" || subUrl == "welcome" || subUrl == "chat" {
		return
	}

	fmt.Println("eyo in here")
	_, ok := ctx.Input.Session("id").(int)
	if !ok {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.WriteString("Unauothorized!")
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, checkForAuthorization)
	beego.Router("welcome", &controllers.MainController{})
	beego.Router("register", &controllers.RegisterController{})
	beego.Router("login", &controllers.LoginController{})
	beego.Router("bot", &controllers.BotController{})
	beego.Router("search", &controllers.ShopsSearchController{})                 // searches for nearest shops
	beego.Router("shops/nearestshop", &controllers.NearestShopItemsController{}) // uses result from nearest shop search to display items
	beego.Router("scan", &controllers.ScanController{})
	beego.Router("buyitem", &controllers.BuyItemController{})

	beego.Router("chat", &controllers.ChatController{})
}
