package controllers

import (
	"encoding/base64"
	"polypanda/apiserver/models"
	"strconv"

	"github.com/astaxie/beego"
)

/*PandaController controller for Panda*/
type PandaController struct {
	beego.Controller
}

/*GetPandas return an array of panda in JSON*/
func (c *PandaController) GetPandas() {
	var ret RetPanda
	pandas, _, err := models.QueryAllPandas()
	if err != nil {
		ret.statusFail(err.Error())
	} else {
		ret.statusOK(pandas)
	}
	c.Ctx.Output.Body(ret.getJSONStream())
}

/*GetByPage return an array of panda in JSON*/
func (c *PandaController) GetByPage() {
	var ret RetPanda
	var num int
	c.Ctx.Input.Bind(&num, "num")
	if num < 1 {
		num = 1
	}
	pandas, _, err := models.QueryPandaByPage(num)
	if err != nil {
		ret.statusFail(err.Error())
	} else {
		ret.statusOK(pandas)
		ret.Extra = strconv.Itoa(num)
	}
	c.Ctx.Output.Body(ret.getJSONStream())
}

/*SetName return an array of panda by page in JSON*/
func (c *PandaController) SetName() {
	var ret RetSimple
	owner := c.Ctx.Input.Param(":owner")
	name := c.Ctx.Input.Param(":ownername")
	decoded, err := base64.StdEncoding.DecodeString(owner)
	if err != nil {
		ret.statusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	err = models.SetNameByOwner(name, decoded)
	if err != nil {
		ret.statusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	ret.statusOK()
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*SetURL return an array of panda by page in JSON*/
func (c *PandaController) SetURL() {
	var ret RetSimple
	idx, err := c.GetInt(":index")
	if err != nil {
		ret.statusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}

	addr := c.Ctx.Input.Param("addr")
	if len(addr) <= 0 {
		ret.statusFail("no photo address ")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	//Encode addr
	err = models.SetURLByIndex(idx, addr)
	if err != nil {
		ret.statusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	ret.statusOK()
	c.Data["json"] = &ret
	c.ServeJSON()
}
