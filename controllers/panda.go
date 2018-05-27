package controllers

import (
	"encoding/base64"
	"log"
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
	pandas, cnt, err := models.QueryAllPandas()
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else {
		ret.StatusOK(pandas)
		ret.Value = strconv.Itoa(cnt)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

func respPandaByPage(c *PandaController, num int) {
	var ret models.RetPanda
	if num < 1 {
		num = 1
	}
	pandas, cnt, err := models.QueryPandaByPage(num)
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else {
		ret.StatusOK(pandas)
		ret.Value = strconv.Itoa(cnt)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*GetPandaByIndex Controller for get panda by id*/
func (c *PandaController) GetPandaByIndex() {
	var ret models.RetPanda
	idx := models.ConvtNumGRTEqZero(c.Ctx.Input.Param(":idx"))
	if idx < 0 {
		ret.StatusFail(models.St400BadRequest)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	pandas, num, err := models.QueryPandaByIndex(idx)
	if err != nil {
		ret.SetStatus(models.St404NotFound, err.Error(), 0)
	} else if num == 0 {
		ret.StatusFail(models.St400BadRequest)
		ret.Value = strconv.Itoa(num)
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

/*SetURL set photo URL of a specific index*/
func (c *PandaController) SetURL() {
	var ret models.RetSimple
	idxStr := c.Ctx.Input.Param(":idx")
	idx, err := strconv.Atoi(idxStr)

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
	log.Println("SetURL idx", idx)
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

/*SetSnapURL set snap URL of a specific index*/
func (c *PandaController) SetSnapURL() {
	var ret models.RetSimple
	idxStr := c.Ctx.Input.Param(":idx")
	idx, err := strconv.Atoi(idxStr)

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
	log.Println("SetSnapURL idx", idx)
	//Encode addr
	err = models.SetSnapURLByIndex(idx, addr)
	if err != nil {
		ret.SetStatus(models.St409Conflict, err.Error(), 0)
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	ret.StatusOK()
	c.Data["json"] = &ret
	c.ServeJSON()
}
