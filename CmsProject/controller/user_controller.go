package controller

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"iris-test-project/CmsProject/service"
	"iris-test-project/CmsProject/utils"
)

type UserController struct {
	Service service.UserService

	Sessions *sessions.Sessions

	Logger *golog.Logger
}

// 获取用户总数
// url:		/v1/users/count
// type: 	get
func (uc *UserController) GetCount() mvc.Result {
	count, err := uc.Service.GetUserTotalCount()
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//正常情况的返回值
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  count,
		},
	}
}
