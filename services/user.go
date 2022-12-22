package services

import (
	"rest_api/infra/helper"
	"rest_api/request"

	"rest_api/models"
	"rest_api/repositories"
)

func Register(input request.AuthenticationInput) (models.User, error) {
	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	savedUser, err := repositories.Save(user)
	if err != nil {
		return user, err
	}
	return savedUser, nil
}

func Login(input request.AuthenticationInput) (string, error) {
	user, err := repositories.FindUserByUserName(input.Username)
	if err != nil {
		return "", err
	}
	err = user.ValidatePassword(input.Password)
	if err != nil {
		return "", err
	}
	jwt, err := helper.GenerateJwt(user)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
