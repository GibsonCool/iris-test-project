package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris-test-project/CmsProject/service"
	"iris-test-project/CmsProject/utils"
)

// 管理员控制器
type AdminController struct {
	Service service.AdminService

	Sessions *sessions.Sessions

	Logger *golog.Logger
}

const (
	ADMINTABLENAME = "admin"
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// 登录功能
// url:   /admin/login
// type:  post
func (ac *AdminController) PostLogin(c iris.Context) mvc.Result {
	ac.Logger.Info(" admin login")
	adminLogin := AdminLogin{}

	// 读取参数
	c.ReadJSON(&adminLogin)

	// 检验参数
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或密码为空,请重新填写后尝试登录",
			},
		}
	}

	// 业务逻辑，数据库查询
	admin, exist := ac.Service.VerifyUser(adminLogin.UserName, adminLogin.Password)
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "登录失败",
				"message": "用户名或者密码错误,请重新登录",
			},
		}
	}

	// 校验通过，设置session
	//userByte, _ := json.Marshal(admin)  内部会判断并且调用 sessions.DefaultTranscoder.Marshal 序列化，所以这里就不用自己序列一次了
	ac.Sessions.Start(c).Set(ADMIN, admin)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}

}

// 获取管理员信息接口
// url:   /admin/info
// type:  get
func (ac *AdminController) GetInfo(c iris.Context) mvc.Result {
	admin := ac.Sessions.Start(c).Get(ADMIN)
	if admin == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.EEROR_UNLOGIN,
				"message": utils.Recode2Text(utils.EEROR_UNLOGIN),
			},
		}
	}
	//解析成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data":   admin,
		},
	}
}

/**
 * 管理员退出功能
 * 请求类型：Get
 * 请求url：admin/singout
 */
func (ac *AdminController) GetSingout(c iris.Context) mvc.Result {

	//删除session，下次需要从新登录
	ac.Sessions.Start(c).Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SIGNOUT),
		},
	}
}

/**
 * 处理获取管理员总数的路由请求
 * 请求类型：Get
 * 请求Url：admin/count
 */
func (ac *AdminController) GetCount() mvc.Result {

	count, err := ac.Service.GetAdminCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERRORADMINCOUNT),
				"count":   0,
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}
}
