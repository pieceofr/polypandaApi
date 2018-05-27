package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"polypanda/polypandaApi/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/status", &controllers.MainController{}, "get:GetStatus")
	//panda Controller
	beego.Router("/api/v1/panda", &controllers.PandaController{})
	beego.Router("/api/v1/panda/:idx([0-9]+)", &controllers.PandaController{}, "get:GetPandaByIndex")
	beego.Router("/api/v1/panda/count", &controllers.PandaController{}, "get:Count")
	beego.Router("/api/v1/panda/:owner/ownername", &controllers.PandaController{}, "put:SetName")
	beego.Router("/api/v1/panda/:idx([0-9]+)/photo", &controllers.PandaController{}, "put:SetURL")
	beego.Router("/api/v1/panda/:idx([0-9]+)/snap", &controllers.PandaController{}, "put:SetSnapURL")
	//Display Controller
	beego.Router("/api/v1/display/pagesize", &controllers.DisplayController{})

}
