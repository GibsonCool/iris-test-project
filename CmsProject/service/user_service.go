package service

import (
	"github.com/go-xorm/xorm"
	"iris-test-project/CmsProject/model"
)

type UserService interface {
	//获取用户日增长统计数据
	GetUserDailyStatisCount(datetime string) int64
	//获取用户总数
	GetUserTotalCount() (int64, error)
	//用户列表
	GetUserList(offset, limit int) []*model.User
}

type userService struct {
	engine *xorm.Engine
}

func NewUserService(engine *xorm.Engine) *userService {
	return &userService{engine: engine}
}

func (u userService) GetUserDailyStatisCount(datetime string) int64 {
	panic("implement me")
}

func (u userService) GetUserTotalCount() (int64, error) {
	//查询del_flag 为0 的用户的总数量；del_flag:0 正常状态；del_flag:1 用户注销或者被删除
	return u.engine.Where("del_flag = 0").Count(model.User{})
}

func (u userService) GetUserList(offset, limit int) []*model.User {
	panic("implement me")
}
