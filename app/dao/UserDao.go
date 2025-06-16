package dao

import (
	"mliev.com/template/go-web/app/Model"
	"mliev.com/template/go-web/support/db"
)

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := db.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
