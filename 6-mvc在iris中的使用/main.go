package main

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var logger *golog.Logger

func main() {
	app := iris.New()
	logger = app.Logger()

	// 注册 MVC 控制器
	//mvc.New(app).Handle(&UserController{})
	mvc.New(app).Handle(new(UserController))

	// 通过 mvc.Configure 配置路由组和控制器
	mvc.Configure(app.Party("/order"), func(mvc *mvc.Application) {
		mvc.Handle(new(OrderController))
	})

	app.Run(iris.Addr(":8000"))
}

func SelfMiddlewareHandler(c iris.Context) {
	logger.Info("调动前" + c.Path())
	c.Next()
	logger.Info("调动后" + c.Path())
}
