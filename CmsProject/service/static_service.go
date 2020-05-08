package service

import (
	"github.com/go-xorm/xorm"
	"iris-test-project/CmsProject/model"
	"iris-test-project/CmsProject/utils"
	"time"
)

type StaticService interface {
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
	GetAdminDailyCount(date string) int64
}

func NewStaticService(engine *xorm.Engine) *staticService {
	return &staticService{engine}
}

type staticService struct {
	Engine *xorm.Engine
}

func (s staticService) InitDate(date string) (string, string, error) {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format(utils.FormatDateStr)
	}

	start, err := time.Parse(utils.FormatDateStr, date)
	if err != nil {
		return "", "", err
	}

	end := start.AddDate(0, 0, 1)

	return start.Format(utils.FormatTimeStr), end.Format(utils.FormatTimeStr), nil
}

// 查询某一日用户增长数量
func (s staticService) GetUserDailyCount(date string) int64 {
	startDate, endDate, err := s.InitDate(date)
	if err != nil {
		return 0
	}

	count, err := s.Engine.
		Where("register_time between ? and ? and del_flag = 0",
			startDate,
			endDate,
		).
		Count(model.User{})
	if err != nil {
		return 0
	}
	return count
}

// 查询某一日订单增长数量
func (s staticService) GetOrderDailyCount(date string) int64 {
	startDate, endDate, err := s.InitDate(date)
	if err != nil {
		return 0
	}
	count, err := s.Engine.
		Where("time between ? and ? and del_flag = 0",
			startDate,
			endDate,
		).
		Count(model.UserOrder{})
	if err != nil {
		return 0
	}
	return count
}

// 查询某一日管理员增长数量
func (s staticService) GetAdminDailyCount(date string) int64 {
	//查询日期date格式解析
	startDate, endDate, err := s.InitDate(date)
	if err != nil {
		return 0
	}

	count, err := s.Engine.
		Where("create_time between ? and ? and status = 0",
			startDate,
			endDate,
		).
		Count(model.Admin{})
	if err != nil {
		return 0
	}
	return count
}
