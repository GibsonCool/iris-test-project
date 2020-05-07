package main

import (
	"fmt"
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

	engine.ShowSQL(true)                     //设置显示SQL语句
	engine.Logger().SetLevel(core.LOG_DEBUG) //设置日志级别

	//SnakeMapper_在每个单词（大写）之间插入（下划线）以获得表名或列名。
	//engine.SetMapper(core.SnakeMapper{})
	//engine.Sync2(new(UserTable))

	//SameMapper 在结构和表之间使用相同的名称。
	//engine.SetMapper(core.SameMapper{})
	//engine.Sync2(new(StudentTable))

	//GonicMapper与SnakeMapper基本相同，但在常用的缩写之间不插入下划线。例如ID，id在GonicMapper中转换为，但在SnakeMapper ID中转换为i_d。
	engine.SetMapper(core.GonicMapper{})
	engine.Sync2(new(PersonTable))

	if tableEmpty, err := engine.IsTableEmpty(new(UserTable)); err != nil {
		panic(err.Error())
	} else {
		if tableEmpty {
			fmt.Println("用户表数据为空")
		} else {
			fmt.Println("用户表数据不为空")
		}
	}

	if tableExist, err := engine.IsTableExist(new(StudentTable)); err != nil {
		panic(err.Error())
	} else {
		if tableExist {
			fmt.Println("学生表存在")
		} else {
			fmt.Println("学生表不存在")
		}

	}
}

type UserTable struct {
	UserId   int64  `xorm:"pk autoincr"`
	UserName string `xorm:"varchar(32)"` //用户名
	UserAge  int64  `xorm:"default 1"`   //用户年龄
	UserSex  int64  `xorm:"default 0"`   //用户性别
}

/**
 * 学生表
 */
type StudentTable struct {
	Id          int64  `xorm:"pk autoincr"` //主键 自增
	StudentName string `xorm:"varchar(24)"` //
	StudentAge  int    `xorm:"int default 0"`
	StudentSex  int    `xorm:"index"` //sex为索引
}

/**
 * 人员结构表
 */
type PersonTable struct {
	Id         int64     `xorm:"pk autoincr"`   //主键自增
	PersonName string    `xorm:"varchar(24)"`   //可变字符
	PersonAge  int       `xorm:"int default 0"` //默认值
	PersonSex  int       `xorm:"notnull"`       //不能为空
	City       CityTable `xorm:"-"`             //不映射该字段，则在数据库表创建的时候就不会有这个字段信息
}

type CityTable struct {
	CityName      string
	CityLongitude float32
	CityLatitude  float32
}
