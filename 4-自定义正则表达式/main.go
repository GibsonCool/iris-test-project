package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"net/http"
	"strconv"
)

func main() {
	app := iris.New()

	// 1.handler 方式处理请求,其实 app.Get 内部也是调用了 app.Handle
	app.Handle(http.MethodGet, "/userInfo", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		app.Logger().Error("request path:" + path)
	})
	// 自定义 Post
	app.Handle(http.MethodPost, "/postCommit", func(c context.Context) {
		path := c.Path()
		app.Logger().Info("post request ,the url path is :" + path)
	})

	// 正则表达式获取路由参数
	// 使用  context.Params().Get("xxx") 可以获取到  url 中 /hello/{xxx} 的参数值、
	app.Get("/hello/{name}", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		// 获取参数
		name := c.Params().Get("name")
		c.HTML("<h1>" + name + "</h1>")
	})

	// 自定义正则表达式，可对变量进行类型限制
	app.Get("/api/id/{userId:uint}", func(c context.Context) {
		userId, err := c.Params().GetInt("userId")
		if err != nil {
			c.JSON(map[string]interface{}{
				"code":    201,
				"message": "参数有误",
			})
			return
		}
		c.JSON(map[string]interface{}{
			"code":    200,
			"message": "userID：" + strconv.Itoa(userId),
		})
	})

	app.Get("/api/users/{isLogin:bool}", func(c context.Context) {
		isLogin, err := c.Params().GetBool("isLogin")
		if err != nil {
			c.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			c.WriteString("已登录")
		} else {
			c.WriteString("未登录")
		}
	})

	app.Get("/api/{name:string}/{isLogin:bool}", func(c context.Context) {
		isLogin, err := c.Params().GetBool("isLogin")
		name := c.Params().GetString("name")
		if err != nil {
			c.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			c.WriteString(name + " 已登录")
		} else {
			c.WriteString(name + " 未登录")
		}
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
