package service

import (
	"github.com/go-xorm/xorm"
	"iris-test-project/CmsProject/model"
)

type OrderService interface {
	GetCount() (int64, error)
}

type orderService struct {
	Engine *xorm.Engine
}

func NewOrderService(engine *xorm.Engine) *orderService {
	return &orderService{Engine: engine}
}

func (o *orderService) GetCount() (int64, error) {
	return o.Engine.Where("del_flag =0").Count(model.UserOrder{})
}
