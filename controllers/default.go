package controllers

import (
	"polypanda/polypandaApi/models"

	"github.com/astaxie/beego"
)

/*MainController Default Controller*/
type MainController struct {
	beego.Controller
}

/*Get Return Default HTML Page*/
func (c *MainController) Get() {
	c.Data["Website"] = "PolyPanda Beego APIs"
	c.Data["Email"] = "polypanda@gmail.com"
	c.TplName = "index.tpl"
}

/*GetStatus keep alive function*/
func (c *MainController) GetStatus() {
	var ret models.RetSimple
	ret.StatusOK()
	c.Data["json"] = &ret
	c.ServeJSON()

}
