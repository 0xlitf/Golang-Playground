package routers

import (
	"github.com/4u1kto/Golang-Playground/BeeProj/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.MainController{})
}
