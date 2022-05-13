package models

import (
	"github.com/jinzhu/gorm"
)

type Wallet struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Balance   string `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetWallets() ([]Wallet, error) {
	var (
		wallets []Wallet
		err     error
	)

	err = db.Find(&wallets).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return wallets, nil
}
