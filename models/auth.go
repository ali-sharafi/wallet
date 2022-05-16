package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func CheckAuth(username string, password string) (int, error) {
	var auth Auth
	err := db.Where(Auth{Username: username}).First(&auth).Error
	if err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password)); err != nil {
		return 0, errors.New("invalid Credentials")
	} else {
		return auth.ID, nil
	}
}
