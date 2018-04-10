package controllers

import (
	"encoding/base64"
	"polypanda/polypandaApi/models"
	"strconv"

	"github.com/astaxie/beego"
)

/*PandaController controller for Panda*/
type PandaController struct {
	beego.Controller
}

/*Get return an list of panda in JSON*/
func (c *PandaController) Get() {
	pageNum := models.ConvtNumGRTZero(c.GetString("page"))
	if pageNum < 0 { // did not specify page
		respPandaAll(c)
	}
	//query by page
	respPandaByPage(c, pageNum)
}

func respPandaAll(c *PandaController) {
	var ret models.RetPanda
	pandas, _, err := models.QueryAllPandas()
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else {
		ret.StatusOK(pandas)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

func respPandaByPage(c *PandaController, num int) {
	var ret models.RetPanda
	if num < 1 {
		num = 1
	}
	pandas, numOfrecords, err := models.QueryPandaByPage(num)
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else {
		ret.StatusOK(pandas)
		ret.Value = strconv.Itoa(numOfrecords)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*GetPandaByIndex Controller for get panda by id*/
func (c *PandaController) GetPandaByIndex() {
	var ret models.RetPanda
	idx := models.ConvtNumGRTEqZero(c.Ctx.Input.Param(":index"))
	if idx < 0 {
		ret.StatusFail(models.St400BadRequest)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	pandas, num, err := models.QueryPandaByIndex(idx)
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else {
		ret.StatusOK(pandas)
		ret.Value = strconv.Itoa(num)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*Count Controller for get panda by id*/
func (c *PandaController) Count() {
	var ret models.RetSimple

	count, err := models.QueryRowCount(beego.AppConfig.String("pandatable"))
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else {
		ret.StatusOK()
		ret.Value = strconv.Itoa(count)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*SetName return an array of panda by page in JSON*/
func (c *PandaController) SetName() {
	var ret models.RetSimple
	owner := c.Ctx.Input.Param(":owner")
	name := c.GetString("name")
	if !models.IsValidOwnername(name) {
		ret.StatusFail(models.St400BadRequest)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	decoded, err := base64.StdEncoding.DecodeString(owner)
	if err != nil {
		ret.SetStatus(models.St400BadRequest, err.Error(), 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	err = models.SetNameByOwner(name, decoded)
	if err != nil {
		ret.SetStatus(models.St409Conflict, err.Error(), 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}

	ret.StatusOK()
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*SetURL return an array of panda by page in JSON*/
func (c *PandaController) SetURL() {
	var ret models.RetSimple
	idx, err := c.GetInt(":index")
	if err != nil {
		ret.SetStatus(models.St400BadRequest, err.Error(), 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	var addr string
	err = c.Ctx.Input.Bind(&addr, "addr")
	if err != nil {
		ret.SetStatus(models.St400BadRequest, err.Error(), 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if len(addr) <= 0 {
		ret.SetStatus(models.St400BadRequest, "no photo address", 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}

	//Encode addr
	err = models.SetURLByIndex(idx, addr)
	if err != nil {
		ret.SetStatus(models.St409Conflict, err.Error(), 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	ret.StatusOK()
	c.Data["json"] = &ret
	c.ServeJSON()
}
