package repositories

import (
	"errors"
	"rest_api/infra/database"
	"rest_api/models"
)

func FindUserByUserName(username string) (models.User, error) {
	var user models.User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return user, err
	}
	if user.Username == "" {
		return user, errors.New("invalid username")
	}
	return user, nil
}

func Save(user models.User) (models.User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
