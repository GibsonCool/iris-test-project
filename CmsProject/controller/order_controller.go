package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris-test-project/CmsProject/service"
	"iris-test-project/CmsProject/utils"
	"math/rand"
)

type OrderController struct {
	Service service.OrderService

	Sessions *sessions.Sessions

	Logger *golog.Logger
}

// 查询订单记录总数
// url： /bos/orders/count
// type: get
func (oc *OrderController) GetCount() mvc.Result {
	oc.Logger.Info(" 查询订单记录总数")
	count, err := oc.Service.GetCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	// TODO试数据返回随机数
	count = rand.Int63n(10)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}
}
