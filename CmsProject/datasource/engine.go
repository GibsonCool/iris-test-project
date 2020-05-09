package datasource

import (
	_ "github.com/go-sql-driver/mysql" // 导入驱动
	"github.com/go-xorm/xorm"
	"iris-test-project/CmsProject/model"
	"xorm.io/core"
)

func NewMysqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/elmcms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	engine.ShowSQL(true)                     //设置显示SQL语句
	engine.Logger().SetLevel(core.LOG_DEBUG) //设置日志级别

	err = engine.Sync2(
		new(model.Permission),
		new(model.City),
		new(model.Admin),
		new(model.AdminPermission),
		new(model.User),
		new(model.UserOrder),
		new(model.Address),
		new(model.Shop),
		new(model.OrderStatus),
		new(model.FoodCategory),
	)

	if err != nil {
		panic(err.Error())
	}

	return engine
}
