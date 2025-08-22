package dao

import (
	"cnb.cool/mliev/examples/go-web/app/model"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
)

// TestDemoDao 注入
type TestDemoDao struct {
	database interfaces.DatabaseInterface
}

func (receiver *TestDemoDao) GetUserByUsername(username string) (*model.TestDemo, error) {
	var user model.TestDemo
	if err := receiver.database.Where("username = ?", username).First(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
