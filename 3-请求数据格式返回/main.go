package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Student struct {
	StuName string `xml:"stu_name"`
	StuAge  int    `xml:"stu_age"`
}

func main() {
	app := iris.New()

	app.Get("/getJson", func(c context.Context) {
		c.JSON(iris.Map{"message": "hello world", "requestCode": 200})
	})

	app.Get("/getStuJson", func(c context.Context) {
		c.JSON(Student{StuName: "double", StuAge: 25})
	})

	app.Get("/getXml", func(c context.Context) {
		c.XML(Student{StuName: "double", StuAge: 44})
	})

	app.Get("/helloText", func(c context.Context) {
		c.Text("text hello  world")
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
