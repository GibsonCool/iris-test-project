package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris-test-project/CmsProject/service"
	"iris-test-project/CmsProject/utils"
	"math/rand"
	"strings"
)

// 统计功能控制器
type StaticController struct {
	Service service.StaticService

	Sessions *sessions.Sessions

	Logger *golog.Logger
}

var (
	ADMINMODULE = "ADMIN_"
	USERMODULE  = "USER_"
	ORDERMODULE = "ORDER_"
)

// 解析路由统计功能请求
// url:  static/{model}/{date}/count
// type: get
func (sc *StaticController) GetCount(c iris.Context) mvc.Result {

	path := c.Path()
	var pathSlice []string
	if path != "" {
		pathSlice = strings.Split(path, "/")
	}

	sc.Logger.Info(path, pathSlice)
	// 不符合我们定义的正则路由匹配路径
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	model := c.Params().Get("model")
	date := c.Params().Get("date")

	var result int64
	session := sc.Sessions.Start(c)
	switch model {
	case "user":
		keyName := USERMODULE + date
		// 先从缓存中取，如果有则直接返回
		result := session.Get(keyName)
		if result != nil {
			result := result.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  result,
				},
			}
		} else {
			// 通过业务层从 DB 查询获取
			result = sc.Service.GetUserDailyCount(date)
			// 将数据设置缓存到 session 中。由于 session 在 main 中使用的 Redis 作为存储，所以数据会直接通过 redis 来缓存
			session.Set(keyName, result)
		}
	case "order":
		keyName := ORDERMODULE + date
		// 先从缓存中取，如果有则直接返回
		result := session.Get(keyName)
		if result != nil {
			result := result.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  result,
				},
			}
		} else {
			result = sc.Service.GetOrderDailyCount(date)
			session.Set(keyName, result)
		}
	case "admin":
		keyName := ADMINMODULE + date
		// 先从缓存中取，如果有则直接返回
		result := session.Get(keyName)
		if result != nil {
			result := result.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  result,
				},
			}
		} else {
			result = sc.Service.GetAdminDailyCount(date)
			session.Set(keyName, result)
		}
	}

	// TODO： 测试数据返回随机数
	result = rand.Int63n(100)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}
}
