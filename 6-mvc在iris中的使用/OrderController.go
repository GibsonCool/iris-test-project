package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

// 处理订单相关的控制器
type OrderController struct{}

// url:  http://localhost:8000/order
// type: get
func (uc *OrderController) Get() string {
	logger.Info("get 请求")
	return "获取订单  get "
}

// url:  http://localhost:8000/order
// type: post
func (uc *OrderController) Post() {
	logger.Info("post 请求  获取订单")
}

// url:  http://localhost:8000/order/info
// type: get
func (uc *OrderController) GetInfo() mvc.Result {
	logger.Info("get 请求,请求路径为 /order/info")
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 1,
			"msg":  "查询订单成",
		},
	}
}

// url:  http://localhost:8000/order/{path:string}/{orderId:long}
// type: get 		多个正则参数的时候，方法参数顺序要匹配，否则找不到路由
func (uc *OrderController) GetBy(path string, orderId int) mvc.Result {
	logger.Info("get 请求,请求路径为 " + path + "/" + strconv.Itoa(orderId))
	return mvc.Response{
		Object: map[string]interface{}{
			"orderId": orderId,
			"name":    "请求订单成功",
			"path":    path,
		},
	}
}

func (uc *OrderController) BeforeActivation(a mvc.BeforeActivation) {
	// http://localhost:8000/order/insert/{id:long}   type:get      UserInsert 方法   不使用中间件
	a.Handle(iris.MethodGet, "/insert/{id:long}", "OrderInsert", SelfMiddlewareHandler)
}

func (uc *OrderController) OrderInsert(id int64) mvc.Result {
	logger.Info("order info insert")
	return mvc.Response{
		Object: map[string]interface{}{
			"msg":    "订单插入成功",
			"userID": id,
		},
	}
}
