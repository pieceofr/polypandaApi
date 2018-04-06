package controllers

import (
	"encoding/base64"
	"fmt"
	"log"
	"polypanda/apiserver/models"
	"strconv"

	"github.com/astaxie/beego"
)

/*PandaController controller for Panda*/
type PandaController struct {
	beego.Controller
}

/*Get return an array of panda in JSON*/
func (c *PandaController) AllPandas() {
	var ret RetPanda
	pandas, _, err := models.QueryAllPandas()
	if err != nil {
		ret.statusFail(err.Error())
	} else {
		ret.statusOK(pandas)
	}
	c.Ctx.Output.Body(ret.getJSONStream())
}

/*PandaByPage return an array of panda in JSON*/
func (c *PandaController) PandaByPage() {
	var ret RetPanda
	num, err := c.GetInt("num")
	if err != nil {
		ret.statusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	if num <= 0 {
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
	owner := c.GetString(":owner")
	name := c.GetString(":name")
	log.Println("owner:", owner, " name:", name)
	decoded, err := base64.StdEncoding.DecodeString(owner)
	if err != nil {
		ret.statusFail(err.Error())
		c.Data["json"] = &ret
		c.ServeJSON()
	}
	fmt.Printf("decoded %s\n", decoded)
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

/*SetNumPage set how number of pandas in a page*/
func (c *PandaController) SetNumPage() {
	var ret RetSimple
	num, err := c.GetInt(":page")
	if err != nil {
		ret.statusFail("")
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

/*GetNumPage set how number of pandas in a page*/
func (c *PandaController) GetNumPage() {
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

/*SetURL return an array of panda by page in JSON*/
func (c *PandaController) SetURL() {
	var ret RetSimple
	idx, err := c.GetInt(":index")
	if err != nil {
		ret.statusFail(err.Error())
	}
	addr := c.GetString("addr")
	log.Println("index:", idx, " addr:", addr)
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
