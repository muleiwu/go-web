package dao

import (
	"cnb.cool/mliev/examples/go-web/app/model"
	"cnb.cool/mliev/examples/go-web/helper"
)

func GetUserByUsername(username string) (*model.TestDemo, error) {
	var user model.TestDemo
	if err := helper.Database().Where("username = ?", username).First(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
