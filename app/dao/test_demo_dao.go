package dao

import (
	"mliev.com/template/go-web/app/model"
	"mliev.com/template/go-web/helper"
)

func GetUserByUsername(username string) (*model.TestDemo, error) {
	var user model.TestDemo
	if err := helper.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
