package script

import (
	"Golang_RPG/models"
	"fmt"
	"net/url"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func convert_datasource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	beego.Info(result)
	return
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	var connectionString string
	if os.Getenv("GO_ENV") == "production" {
		connectionString = convert_datasource(os.Getenv("DATABASE_URL"))
	} else {
		connectionString = beego.AppConfig.String("connectionString")
	}

	maxIdle := 0
	maxConn := 30

	orm.RegisterDataBase("default", "mysql", connectionString, maxIdle, maxConn)
	orm.RegisterModel(new(models.Users))
	orm.RegisterModel(new(models.Bots))
	orm.RegisterModel(new(models.Inventory))
	orm.RegisterModel(new(models.Enemies))
	orm.RegisterModel(new(models.Locations))
	orm.RegisterModel(new(models.ShopItems), new(models.Items))
}

func init() {
	fmt.Println("hello")
}
