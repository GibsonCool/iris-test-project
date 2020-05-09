package datasource

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"iris-test-project/CmsProject/config"
)

func NewRedis() *redis.Database {
	var dataBase *redis.Database

	appConfig := config.InitConfig()
	if appConfig != nil {
		rd := appConfig.Redis
		dataBase = redis.New(service.Config{
			Network:     rd.NetWork,
			Addr:        rd.Addr + ":" + rd.Port,
			Password:    rd.Password,
			Database:    "",
			MaxIdle:     0,
			MaxActive:   10,
			IdleTimeout: service.DefaultRedisIdleTimeout,
			Prefix:      rd.Prefix,
		})
	} else {
		iris.New().Logger().Info("Redis  创建失败")
	}

	return dataBase
}
