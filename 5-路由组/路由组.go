package main

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var logger *golog.Logger

func main() {
	app := iris.New()
	logger = app.Logger()
	// 业务经常需要按模块分组
	// 例如用户模块下的请求有 xxx/user/register   xxx/user/login    xxx/user/info

	// 路由组请求，以及中间件使用，
	userParty := app.Party("/user", func(c context.Context) {
		app.Logger().Error("user 的中间件调用前")
		// 一定要在调用 next 之后，后续的路由例如 /user/register 才会被执行到
		c.Next()
		app.Logger().Error("user 的中间件调用后")
	}, selfMiddlewareHandler)

	// 请求路径为： /user/register
	userParty.Get("/register", func(c context.Context) {
		app.Logger().Info("用户注册功能")
		c.HTML("<h1>用户注册功能</h1>")
	})
	// 请求路径为： /user/login
	userParty.Get("/login", func(c context.Context) {
		app.Logger().Info("用户登录功能")
		c.HTML("<h1>用户登录功能</h1>")
	})

	adminParty := app.Party("/admin")

	// 中间件也可以这样用
	adminParty.Use(selfMiddlewareHandler)

	// 和 Use() 不同，Done() 这里的中间件，总是在路由请求例如：/admin/info 处理完之后再调用 c.next() 之后才会执行
	adminParty.Done(func(c context.Context) {
		app.Logger().Info("response send to " + c.Path())
		// 这里调用 next() 才会触发 doneMiddlewareHandler 中间件的执行，不掉用，则不会执行
		c.Next()
	}, doneMiddlewareHandler)

	// 请求路径为： /admin/info
	adminParty.Get("/info", func(c context.Context) {
		c.HTML("<h1>用户信息</h1>")
		// 这里调用 next() 触发 done() 中的中间件
		c.Next()
	})

	// 请求路径为： /admin/query
	adminParty.Get("/query", func(c context.Context) {
		c.HTML("<h1>查询信息</h1>")
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

func selfMiddlewareHandler(ctx context.Context) {
	logger.Info("自定义中间件调用前：" + ctx.Path())
	ctx.Next()
	logger.Info("自定义中间件调用后：" + ctx.Path())
}

func doneMiddlewareHandler(ctx context.Context) {
	logger.Info("done中间件调用前：" + ctx.Path())
	ctx.Next()
	logger.Info("done中间件调用后：" + ctx.Path())
}
