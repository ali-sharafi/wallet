package utils

import (
	"errors"

	"github.com/ali-sharafi/wallet/pkg/settings"
	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.AppSetting.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("unAuthorized")
	}
}
