package main

import "github.com/kataras/iris"

// iris 配置设置案例
func main() {
	app := iris.New()

	// iris 的配置参数可以参考一下 https://www.studyiris.com/doc/irisDoc/Configuration.html

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
