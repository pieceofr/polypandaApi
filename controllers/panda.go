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
	var ret models.RetPanda
	pandas, _, err := models.QueryAllPandas()
	if err != nil {
		ret.StatusFail(err.Error())
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
	pandas, _, err := models.QueryPandaByPage(num)
	if err != nil {
		ret.StatusFail(err.Error())
	} else {
		ret.StatusOK(pandas)
		ret.Extra = strconv.Itoa(num)
	}
	c.Data["json"] = &ret
	c.ServeJSON()
}

/*SetName return an array of panda by page in JSON*/
func (c *PandaController) SetName() {
	var ret models.RetSimple
	owner := c.Ctx.Input.Param(":owner")
	name := c.Ctx.Input.Param(":ownername")
	decoded, err := base64.StdEncoding.DecodeString(owner)
	if err != nil {
		ret.StatusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	err = models.SetNameByOwner(name, decoded)
	if err != nil {
		ret.StatusFail(err.Error())
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
		ret.StatusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	var addr string
	err = c.Ctx.Input.Bind(&addr, "addr")
	if err != nil {
		ret.StatusFail("no photo address ")
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if len(addr) <= 0 {
		ret.StatusFail("no photo address ")
		c.Data["json"] = &ret
		c.ServeJSON()
	}

	//Encode addr
	err = models.SetURLByIndex(idx, addr)
	if err != nil {
		ret.StatusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	ret.StatusOK()
	c.Data["json"] = &ret
	c.ServeJSON()
}
