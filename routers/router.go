package routers

import (
	"polypanda/apiserver/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/allPandas", &controllers.PandaController{}, "get:AllPandas")
	beego.Router("/api/v1/setName/:owner/:name", &controllers.PandaController{}, "put:SetName")
	beego.Router("/api/v1/panda/page?:num", &controllers.PandaController{}, "get:PandaByPage")
	beego.Router("/api/v1/numRecordsInPage", &controllers.PandaController{}, "get:GetNumPage")
	beego.Router("/api/v1/numRecordsInPage/:page", &controllers.PandaController{}, "put:SetNumPage")
	beego.Router("/api/v1/:index/url?:addr", &controllers.PandaController{}, "put:SetURL")
}
