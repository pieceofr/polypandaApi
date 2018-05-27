package main

import (
	"polypanda/polypandaApi/models"
	_ "polypanda/polypandaApi/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.SQLConnect()
	defer models.SQLDisconnect()
	beego.Run()
}
