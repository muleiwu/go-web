package dao

import (
	"cnb.cool/mliev/examples/go-web/app/model"
	"cnb.cool/mliev/examples/go-web/helper/database"
)

func GetUserByUsername(username string) (*model.TestDemo, error) {
	var user model.TestDemo
	if err := database.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
