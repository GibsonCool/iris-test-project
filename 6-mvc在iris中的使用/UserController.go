package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
)

// 处理用户相关的控制器
type UserController struct{}

// url:  http://localhost:8000
// type: get
func (uc *UserController) Get() string {
	logger.Info("get 请求")
	return "hello world"
}

// url:  http://localhost:8000
// type: post
func (uc *UserController) Post() {
	logger.Info("post 请求")
}

// url:  http://localhost:8000/info
// type: get
func (uc *UserController) GetInfo() mvc.Result {
	logger.Info("get 请求,请求路径为 info")
	return mvc.Response{
		Object: map[string]interface{}{
			"code": 1,
			"msg":  "请求成功",
		},
	}
}

// 匹配正则表达式路由
// url:  http://localhost:8000/info/{code:long}
// type: get
func (uc *UserController) GetInfoBy(code int) mvc.Result {
	logger.Info("get 请求,请求路径为 info/" + strconv.Itoa(code))
	return mvc.Response{
		Object: map[string]interface{}{
			"code": code,
			"msg":  "请求成功",
		},
	}
}

// 多参数路由
// url:  http://localhost:8000/user/{age:long}/{name:string}
// type: get
func (uc *UserController) GetUserBy(age int, name string) mvc.Result {
	logger.Info("get 请求,请求路径为 user/" + strconv.Itoa(age) + "/" + name)
	return mvc.Response{
		Object: map[string]interface{}{
			"age":  age,
			"name": "请求成功:" + name,
		},
	}
}

// url:  http://localhost:8000/{path:string}/{age:long}/{name:string}
// type: get 		多个正则参数的时候，方法参数顺序要匹配，否则找不到路由
func (uc *UserController) GetBy(path string, age int, name string) mvc.Result {
	logger.Info("get 请求,请求路径为 " + path + "/" + strconv.Itoa(age) + "/" + name)
	return mvc.Response{
		Object: map[string]interface{}{
			"age":  age,
			"name": "请求成功:" + name,
			"path": path,
		},
	}
}

// 不使用 GetXxx PostXx 之类的方法，也可以通过此方法自定义要使用那个方法处理请求
func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
	// http://localhost:8000/query   type:get      UserInfo 方法    SelfMiddlewareHandler 中间件
	a.Handle(iris.MethodGet, "/query", "UserInfo", SelfMiddlewareHandler)

	// http://localhost:8000/insert/{id:long}   type:get      UserInsert 方法   不使用中间件
	a.Handle(iris.MethodGet, "/insert/{id:long}", "UserInsert")

	// http://localhost:8000/delete/{id:long}/{name:string}  type:post      DeleteUserInfo 方法   不使用中间件
	a.Handle(iris.MethodPost, "/delete/{id:long}/{name:string}", "DeleteUserInfo")
}

func (uc *UserController) UserInfo() mvc.Result {
	logger.Info("user info query")
	return mvc.Response{
		Object: map[string]interface{}{
			"name": "double",
			"age":  23,
		},
	}
}

func (uc *UserController) UserInsert(id int64) mvc.Result {
	logger.Info("user info insert")
	return mvc.Response{
		Object: map[string]interface{}{
			"msg":    "插入成功",
			"userID": id,
		},
	}
}

// 方法参数顺序和参数类型与多路由正则参数保持一致
func (uc *UserController) DeleteUserInfo(id int64, name string) mvc.Result {
	logger.Info("user info delete")
	return mvc.Response{
		Object: map[string]interface{}{
			"msg":      "删除成功",
			"userID":   id,
			"userName": name,
		},
	}
}
