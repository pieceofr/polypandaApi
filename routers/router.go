package routers

import (
	"polypanda/polypandaApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//panda Controller
	beego.Router("/api/v1/panda", &controllers.PandaController{})
	beego.Router("/api/v1/panda/:owner/:ownername([\\w-]+)", &controllers.PandaController{}, "put:SetName")
	beego.Router("/api/v1/panda/:index/photo", &controllers.PandaController{}, "put:SetURL")
	//Display Controller
	beego.Router("/api/v1/display/numpage", &controllers.DisplayController{})

}

/* API Design
-  /panda/all (GET) Return all Pandas
-  /panda/page (GET) Return Pandas By Page
	param : num = 1
-  /panda/:owner/:ownername (PUT) Set Panda Ownername
-  /panda/:index/photo (PUT) Assign owner's photo addr
	param : addr = locationOfPhoto
-  /display/numpage (GET) Return number of page to return
-  /display/numpage (PUT) Return number of page to return
	param : num = {number of records in one page}
*/
