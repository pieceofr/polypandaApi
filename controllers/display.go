package controllers

import (
	"polypanda/polypandaApi/models"
	"strconv"

	"github.com/astaxie/beego"
)

/*DisplayController controller for Panda*/
type DisplayController struct {
	beego.Controller
}

/*Get set how number of pandas in a page*/
func (c *DisplayController) Get() {
	var ret models.RetSimple
	_, err := strconv.Atoi(beego.AppConfig.String("numRecordsInPage"))
	if err == nil {
		ret.StatusOK()
		ret.Extra = beego.AppConfig.String("numRecordsInPage")
	} else {
		ret.StatusFail(err.Error())
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*Put set how number of pandas in a page*/
func (c *DisplayController) Put() {
	var ret models.RetSimple
	var num int
	err := c.Ctx.Input.Bind(&num, "num")
	if err != nil {
		ret.StatusFail("")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if num <= 0 {
		ret.StatusFail("Invalid input of  number. number must be a digit and greater than zero")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if models.SetNumInPage(num) {
		ret.StatusOK()

	} else {
		ret.StatusFail("")
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}
