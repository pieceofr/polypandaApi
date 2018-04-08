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
	//pageSize := models.ConvertPageNum(c.GetString("pageSize"))
	pageNum := models.ConvertPageNum(c.GetString("pageNum"))
	log.Println("pageNum:", pageNum)
	if pageNum < 0 {
		respPandaAll(c)
	}
	//query by page
	respPandaByPage(c, pageNum)
}

func respPandaAll(c *PandaController) {
	var ret RetPanda
	pandas, _, err := models.QueryAllPandas()
	if err != nil {
		ret.statusFail(err.Error())
	} else {
		ret.statusOK(pandas)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

func respPandaByPage(c *PandaController, num int) {
	var ret RetPanda
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
	c.Data["json"] = &ret
	c.ServeJSON()
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
	var addr string
	err = c.Ctx.Input.Bind(&addr, "addr")
	if err != nil {
		ret.statusFail("no photo address ")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
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
