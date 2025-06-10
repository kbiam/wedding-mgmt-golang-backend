package auth

import "gorm/db"

func GetUserByUsername(username string)(*AdminUser, error) {
	var user AdminUser
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}