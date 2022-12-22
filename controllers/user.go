package controllers

import (
	"net/http"
	"rest_api/request"
	"rest_api/services"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input request.AuthenticationInput
	err := context.ShouldBindJSON(&input)
	// TODO: create common middleware to handle all validations
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savedUser, err := services.Register(input)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": savedUser})

}

func Login(context *gin.Context) {
	var input request.AuthenticationInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := services.Login(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": jwt})
}
