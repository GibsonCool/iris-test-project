package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iris-test-project/CmsProject/config"
	"iris-test-project/CmsProject/model"
)

func main() {
	app := newApp()

	// 应用配置
	configuration(app)

	// 路由设置
	mvcHandler(app)

	// 初始化配置文件
	appConfig := config.InitConfig()

	// 启动服务
	app.Run(
		iris.Addr(":"+appConfig.Port),
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //让程序自身尽可能的优化
	)
}

func newApp() *iris.Application {
	app := iris.New()

	// 设置日志级别，开发阶段为 debug
	app.Logger().SetLevel("debug")

	// 注册静态资源
	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manage/static", "./static")

	// 注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(c context.Context) {
		c.View("index.html")
	})
	return app
}

func configuration(app *iris.Application) {
	// 配置字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	// 错误码配置处理
	//404
	app.OnErrorCode(iris.StatusNotFound, func(c context.Context) {
		c.JSON(model.BaseResponse{
			ErrMsg: iris.StatusNotFound,
			Msg:    " not found",
			Data:   iris.Map{},
		})
	})

	// 500
	app.OnErrorCode(iris.StatusInternalServerError, func(c context.Context) {
		c.JSON(model.BaseResponse{
			ErrMsg: iris.StatusInternalServerError,
			Msg:    " internal error ",
			Data:   iris.Map{},
		})
	})
}

func mvcHandler(app *iris.Application) {

}
