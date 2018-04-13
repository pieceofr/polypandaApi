package controllers

import (
	"bytes"
	"os/exec"
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
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	//out, err := exec.Command("cat README.md").Output()
	if err != nil {
		ret.Message = err.Error()
		ret.Value = ""
	}
	ret.StatusOK()
	ret.Value = string(bytes.TrimRight(out, "\x0a"))
	c.Data["json"] = &ret
	c.ServeJSON()

}
