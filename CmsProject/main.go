package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris-test-project/CmsProject/config"
	"iris-test-project/CmsProject/controller"
	"iris-test-project/CmsProject/datasource"
	"iris-test-project/CmsProject/model"
	"iris-test-project/CmsProject/service"
	"time"
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
	app.StaticWeb("/img", "./static/img")
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
			Status:  iris.StatusNotFound,
			Message: " not found",
			Data:    iris.Map{},
		})
	})

	// 500
	app.OnErrorCode(iris.StatusInternalServerError, func(c context.Context) {
		c.JSON(model.BaseResponse{
			Status:  iris.StatusInternalServerError,
			Message: " internal error ",
			Data:    iris.Map{},
		})
	})
}

func mvcHandler(app *iris.Application) {

	// 创建 session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessionCookie",
		Expires: 1 * time.Hour,
	})

	sqlEngine := datasource.NewMysqlEngine()

	//管理员模板功能
	adminService := service.NewAdminService(sqlEngine)
	adminGroup := mvc.New(app.Party("/admin"))

	// 两种方式将控制器需要的参数传入
	//adminGroup.Handle(&controller.AdminController{Service: adminService, Sessions: sessManager, Logger: app.Logger()})
	adminGroup.Register(
		adminService,
		sessManager,
		app.Logger(),
	)
	adminGroup.Handle(new(controller.AdminController))

	// 统计功能模块
	staticService := service.NewStaticService(sqlEngine)
	staticGroup := mvc.New(app.Party("/statis/{model}/{date}/"))
	staticGroup.Register(
		staticService,
		sessManager,
		app.Logger(),
	)
	staticGroup.Handle(new(controller.StaticController))

	// 订单功能模块
	orderService := service.NewOrderService(sqlEngine)
	orderGroup := mvc.New(app.Party("/bos/orders/"))
	orderGroup.Register(
		orderService,
		sessManager,
		app.Logger(),
	)
	orderGroup.Handle(new(controller.OrderController))

	// 用户功能模块
	userService := service.NewUserService(sqlEngine)
	userGroup := mvc.New(app.Party("/v1/users/"))
	userGroup.Register(
		userService,
		sessManager,
		app.Logger(),
	)
	userGroup.Handle(new(controller.UserController))
}
