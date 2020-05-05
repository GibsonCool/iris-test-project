package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	app.Get("/getRequest", func(ctx context.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
	})

	// 1.处理 Get 请求
	app.Get("/userpath", func(c context.Context) {
		// 获取路径
		path := c.Path()
		app.Logger().Info(path)
		// 写入返回数据：string类型的
		c.WriteString("请求路径为：" + path)
	})

	// 2.处理Get 请求，并接受参数
	app.Get("/userInfo", func(c context.Context) {
		// 获取路径
		path := c.Path()
		app.Logger().Info(path)

		// 获取 get 请求所携带的参数
		userName := c.URLParam("username")
		app.Logger().Info(userName)

		pwd := c.URLParam("pwd")
		app.Logger().Info(pwd)

		// 写入返回数据：HTML 类型格式
		c.HTML("<h1>" + userName + "," + pwd + "<h1>")
	})

	// 3.处理 Post 请求 form格式请求数据
	app.Post("/postLogin", func(c context.Context) {
		// 获取路径
		path := c.Path()
		app.Logger().Info(path)

		name := c.PostValue("name")
		value := c.PostValue("pwd")
		app.Logger().Info(name, "---", value)

		c.HTML("name：" + name + "   pwd：" + value)
	})

	// 4.处理 Post 请求  json格式请求数据
	app.Post("/postJson", func(c context.Context) {
		// 获取路径
		path := c.Path()
		app.Logger().Info(path)

		var person Person
		if err := c.ReadJSON(&person); err != nil {
			panic(err.Error())
		}

		c.Writef("接收到了json 数据：%#+v\n", person)
	})

	// 4.处理 Post 请求  Xml 格式请求数据
	// Content-type   application/xml
	//<student>
	//		<stu_name>doublex</stu_name>
	//		<stu_age>33</stu_age>
	//</student>
	app.Post("/postXml", func(c context.Context) {
		// 获取路径
		path := c.Path()
		app.Logger().Info(path)

		var student Student
		if err := c.ReadXML(&student); err != nil {
			panic(err.Error())
		}

		c.Writef("接收到了Xml 数据：%#+v\n", student)
	})

	// put 请求
	app.Put("/putInfo", func(c context.Context) {
		path := c.Path()
		app.Logger().Info("请求URL：", path)
	})

	// delete 请求
	app.Delete("/deleteUser", func(c context.Context) {
		path := c.Path()
		app.Logger().Info("Delete 请求 URL：", path)
	})

	app.Run(iris.Addr(":8000"))
}

type (
	Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	Student struct {
		StuName string `xml:"stu_name"`
		StuAge  string `xml:"stu_age"`
	}
)
