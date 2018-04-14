package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Ctx.WriteString("helloworld")
	fmt.Println("header: %s", c.Ctx.Request.Header)
	for k, v := range c.Ctx.Request.Header {
		fmt.Println("%s = %s",k,v)
	}
	fmt.Println("Request: ", *c.Ctx.Request)

	//for k, v := range c.Ctx.Request {
	//	fmt.Println("%s = %s",k,v)
	//}
}
