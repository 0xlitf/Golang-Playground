package routers

import (
	"github.com/4u1kto/Golang-Playground/BeeProj/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/test", &controllers.MainController{})

	//基础路由
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	beego.Any("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	/*
	beego.Get(router, beego.FilterFunc)
	beego.Post(router, beego.FilterFunc)
	beego.Put(router, beego.FilterFunc)
	beego.Patch(router, beego.FilterFunc)
	beego.Head(router, beego.FilterFunc)
	beego.Options(router, beego.FilterFunc)
	beego.Delete(router, beego.FilterFunc)
	beego.Any(router, beego.FilterFunc)
	*/

	//RESTful Controller 路由
	//固定路由
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/admin", &admin.UserController{})
	//beego.Router("/admin/index", &admin.ArticleController{})
	//beego.Router("/admin/addpkg", &admin.AddController{})
	//正则路由
	beego.Router("/api/?:id", &controllers.RController{})

	//默认匹配 //匹配 /api/123 :id = 123 可以匹配 /api/ 这个URL

	//beego.Router("/api/:id”, &controllers.RController{})

	//默认匹配 //匹配 /api/123 :id = 123 不可以匹配 /api/ 这个URL

	beego.Router("/api/:id([0-9]+)", &controllers.RController{})

	//自定义正则匹配 //匹配 /api/123 :id = 123

	beego.Router("/user/:username([\\w]+)", &controllers.RController{})

	//正则字符串匹配 //匹配 /user/astaxie :username = astaxie

	//beego.Router("/download/*.*”, &controllers.RController{})

	//*匹配方式 //匹配 /download/file/api.xml :path= file/api :ext=xml

	beego.Router("/download/ceshi/*", &controllers.RController{})

	//*全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json

	//beego.Router("/:id:int”, &controllers.RController{})

	//int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)

	//beego.Router("/:hi:string”, &controllers.RController{})

	//string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)

	//beego.Router("/cms_:id([0-9]+).html”, &controllers.CmsController{})

	//带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123

	//自定义方法及 RESTful 规则
	//beego.Router("/", &controllers.IndexController{}, "*:Index")
	//使用第三个参数，第三个参数就是用来设置对应 method 到函数名，定义如下
	//
	//*表示任意的 method 都执行该函数
	//使用 httpmethod:funcname 格式来展示
	//多个不同的格式使用 ; 分割
	//多个 method 对应同一个 funcname，method 之间通过 , 来分割
	//以下是一个 RESTful 的设计示例：

	beego.Router("/api/list", &controllers.RestController{}, "*:ListFood")
	beego.Router("/api/create", &controllers.RestController{}, "post:CreateFood")
	beego.Router("/api/update", &controllers.RestController{}, "put:UpdateFood")
	beego.Router("/api/delete", &controllers.RestController{}, "delete:DeleteFood")
	//以下是多个 HTTP Method 指向同一个函数的示例：

	beego.Router("/api", &controllers.RestController{}, "get,post:ApiFunc")
	//以下是不同的 method 对应不同的函数，通过 ; 进行分割的示例：
	//	可用的 HTTP Method：
	//
	//	*: 包含以下所有的函数
	//get: GET 请求
	//post: POST 请求
	//put: PUT 请求
	//delete: DELETE 请求
	//patch: PATCH 请求
	//options: OPTIONS 请求
	//head: HEAD 请求
	//	如果同时存在 * 和对应的 HTTP Method，那么优先执行 HTTP Method 的方法，例如同时注册了如下所示的路由：
	//
	beego.Router("/simple", &controllers.SimpleController{}, "*:AllFunc;post:PostFunc")
	//	那么执行 POST 请求的时候，执行 PostFunc 而不执行 AllFunc。
	//
	//	自定义函数的路由默认不支持 RESTful 的方法，也就是如果你设置了 beego.Router("/api",&RestController{},"post:ApiFunc") 这样的路由，如果请求的方法是 POST，那么不会默认去执行 Post 函数。

	beego.Router("/simple", &controllers.SimpleController{}, "get:GetFunc;post:PostFunc")


	//自动匹配
	
}
