package main

import (
	"polypanda/apiserver/models"
	_ "polypanda/apiserver/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.SQLConnect()
	defer models.SQLDisconnect()
	beego.Run()
}
