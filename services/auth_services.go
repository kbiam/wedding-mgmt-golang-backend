package services

import (
	"errors"
	"gorm/models/entity/auth"
	auth_request "gorm/models/request/auth"
	"gorm/utils"
)

func Login(auth_request auth_request.LoginBody)(*auth.AdminUser, string, error){
	user, err := auth.GetUserByUsername(auth_request.Username)

	if err != nil {
		return nil, "", errors.New("invalid username or password")
	}

	if user.Password != auth_request.Password {
		return nil, "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate JWT token")
	}
	return user, token, nil
}