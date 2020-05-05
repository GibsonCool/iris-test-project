package main

import "github.com/kataras/iris"

func main() {
	// 1.创建 APP结构体对象
	app := iris.New()

	// 2.端口监听
	app.Run(iris.Addr(":7999"), iris.WithoutServerError(iris.ErrServerClosed))
}
