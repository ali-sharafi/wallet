package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ali-sharafi/wallet/models"
	"github.com/ali-sharafi/wallet/pkg/settings"
	"github.com/ali-sharafi/wallet/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	validator "gopkg.in/validator.v2"
)

func GetAuth(c *gin.Context) {
	res := utils.Gin{C: c}
	var user models.Auth
	if err := c.ShouldBindJSON(&user); err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validator.Validate(user); err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	userId, err := models.CheckAuth(user.Username, user.Password)

	if err != nil {
		res.Response(http.StatusUnauthorized, err.Error(), nil)
		return
	}

	if userId == 0 {
		res.Response(http.StatusUnauthorized, "Invalid Credentials", nil)
		return
	}

	token, err := GenerateToken(userId)
	if err != nil {
		res.Response(http.StatusInternalServerError, "Could not login", nil)
		return
	}

	res.Response(http.StatusOK, "Success", token)
}

func GenerateToken(userId int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(userId),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(settings.AppSetting.JwtSecret))

	return token, err
}
