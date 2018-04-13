package routers

import (
	"polypanda/polypandaApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/status", &controllers.MainController{}, "get:GetStatus")
	//panda Controller
	beego.Router("/api/v1/panda", &controllers.PandaController{})
	beego.Router("/api/v1/panda/:idx", &controllers.PandaController{}, "get:GetPandaByIndex")
	beego.Router("/api/v1/panda/count", &controllers.PandaController{}, "get:Count")
	beego.Router("/api/v1/panda/:owner/ownername", &controllers.PandaController{}, "put:SetName")
	beego.Router("/api/v1/panda/:index/photo", &controllers.PandaController{}, "put:SetURL")
	//Display Controller
	beego.Router("/api/v1/display/pagesize", &controllers.DisplayController{})

}
