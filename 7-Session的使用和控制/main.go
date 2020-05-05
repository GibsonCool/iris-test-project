package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
	"strconv"
)

var (
	sessionId = "mySession"
	isLogin   = "IsLogin"
)

func main() {
	app := iris.New()

	// 1.创建 session 并进行使用

	sess := sessions.New(sessions.Config{
		Cookie: sessionId,
	})

	app.Post("/login", func(c context.Context) {
		name := c.PostValue("name")
		pwd := c.PostValue("pwd")

		if name == "double" && pwd == "pwd123" {
			session := sess.Start(c)
			session.Set(isLogin, true)
			c.WriteString("登录成功")
		} else {
			session := sess.Start(c)
			session.Set(isLogin, false)
			c.WriteString("登录失败~~")
		}
	})

	app.Get("/logout", func(c context.Context) {
		session := sess.Start(c)
		session.Set(isLogin, false)
		c.WriteString("退出成功~~")
	})

	app.Get("/query", func(c context.Context) {
		session := sess.Start(c)
		if boolean, err := session.GetBoolean(isLogin); err != nil {
			c.WriteString("请先登录失败~~" + err.Error())
		} else if boolean {
			app.Logger().Info("登录校验成功", boolean)
			c.WriteString("查询成功~~")
		} else {
			c.WriteString("请先登录失败~~" + strconv.FormatBool(boolean))
		}
	})

	// 2.由于 session 是存放在内存中的，为防止程序重启，宕机导致内存数据丢失，还可以让其与后端存储绑定
	// 进行序列化存储，并且可以在重启的时候从新加载
	db, err := boltdb.New("session.db", 0600)
	if err != nil {
		panic(err.Error())
	}
	sess.UseDatabase(db)

	// 注册回调，在服务停止的时候关闭数据库
	// RegisterOnInterrupt注册一个全局函数以在按下CTRL + C / CMD + C或接收到unix kill命令时调用。
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
