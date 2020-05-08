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
	switch model {
	case "user":
		result = sc.Service.GetUserDailyCount(date)
	case "order":
		result = sc.Service.GetOrderDailyCount(date)
	case "admin":
		result = sc.Service.GetAdminDailyCount(date)
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
