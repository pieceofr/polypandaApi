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
	_, err := strconv.Atoi(beego.AppConfig.String("pagesize"))
	if err == nil {
		ret.StatusOK()
		ret.Extra = beego.AppConfig.String("pagesize")
	} else {
		ret.SetStatus(models.St409Conflict, err.Error(), 0)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*Put set how number of pandas in a page*/
func (c *DisplayController) Put() {
	var ret models.RetSimple
	num, err := c.GetInt("size")

	if err != nil {
		ret.StatusFail(models.St400BadRequest)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if num <= 0 {
		ret.SetStatus(models.St400BadRequest, "page size must greater than zero", 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if models.SetNumInPage(num) {
		ret.StatusOK()
		ret.Extra = beego.AppConfig.String("pagesize")
	} else {
		ret.SetStatus(models.St408RequestTimeout, "page size has not changed", 0)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}
