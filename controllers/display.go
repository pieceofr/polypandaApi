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
	var ret RetSimple
	_, err := strconv.Atoi(beego.AppConfig.String("numRecordsInPage"))
	if err == nil {
		ret.statusOK()
		ret.Extra = beego.AppConfig.String("numRecordsInPage")
	} else {
		ret.statusFail(err.Error())
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*Put set how number of pandas in a page*/
func (c *DisplayController) Put() {
	var ret RetSimple
	var num int
	err := c.Ctx.Input.Bind(&num, "num")
	if err != nil {
		ret.statusFail("")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if num <= 0 {
		ret.statusFail("Invalid input of  number. number must be a digit and greater than zero")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if models.SetNumInPage(num) {
		ret.statusOK()

	} else {
		ret.statusFail("")
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}
