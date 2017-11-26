package script

import (
	"Golang_RPG/models"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	var connectionString string
	if os.Getenv("GO_ENV") == "production" {
		connectionString = os.Getenv("DATABASE_URL")
	} else {
		connectionString = beego.AppConfig.String("connectionString")
	}

	orm.RegisterDataBase("default", "mysql", connectionString)
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Bots))
	orm.RegisterModel(new(models.Inventory))
	orm.RegisterModel(new(models.Enemies))
	orm.RegisterModel(new(models.Locations))
	orm.RegisterModel(new(models.ShopItems), new(models.Items))
}
