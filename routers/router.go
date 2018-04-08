package routers

import (
	"polypanda/polypandaApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//panda Controller
	beego.Router("/api/v1/panda", &controllers.PandaController{})
	beego.Router("/api/v1/panda/:owner/ownername", &controllers.PandaController{}, "put:SetName")
	beego.Router("/api/v1/panda/:index/photo", &controllers.PandaController{}, "put:SetURL")
	//Display Controller
	beego.Router("/api/v1/display/pagesize", &controllers.DisplayController{})

}
