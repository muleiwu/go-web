package dao

import (
	"cnb.cool/mliev/open/go-web/app/model"
	"cnb.cool/mliev/open/go-web/pkg/helper"
)

// TestDemoDao 数据访问对象
type TestDemoDao struct {
}

func (receiver *TestDemoDao) GetUserByUsername(username string) (*model.TestDemo, error) {
	var user model.TestDemo
	if err := helper.GetDatabase().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
