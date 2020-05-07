package main

import (
	_ "github.com/go-sql-driver/mysql" // 导入驱动
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	//"用户名:密码@/数据库名称?charset=utf8"
	engine, err := xorm.NewEngine("mysql", "root:123456@/elmcms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	defer engine.Close()

	// 数据库引擎设置
	engine.ShowSQL(true)                     //设置显示SQL语句
	engine.Logger().SetLevel(core.LOG_DEBUG) //设置日志级别
	engine.SetMaxOpenConns(10)               //设置最大连接数
	//engine.SetMaxIdleConns(2)
	engine.Sync2(new(Person))
}

type Person struct {
	Age  int
	Name string
}
