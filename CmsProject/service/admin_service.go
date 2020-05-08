package service

import (
	"github.com/go-xorm/xorm"
	"iris-test-project/CmsProject/model"
)

// 业务操作使用接口定义，便于扩展
type AdminService interface {
	// 通过用户名密码校验登录用户是否合法
	VerifyUser(userName, pwd string) (model.Admin, bool)
	// 获取管理员总数
	GetAdminCount() (int64, error)
}

func NewAdminService(db *xorm.Engine) AdminService {
	return &adminService{engine: db}
}

type adminService struct {
	engine *xorm.Engine
}

func (a *adminService) VerifyUser(userName, pwd string) (model.Admin, bool) {
	var admin model.Admin
	a.engine.Where("admin_name = ? and pwd = ?", userName, pwd).Get(&admin)
	return admin, admin.AdminId != 0
}

func (a *adminService) GetAdminCount() (int64, error) {
	count, err := a.engine.Count(new(model.Admin))
	if err != nil {
		panic(err.Error())
		return 0, err
	}
	return count, nil
}
