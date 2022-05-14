package models

import (
	"errors"
	"strconv"
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

	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func GetBalance(walletID int) (*string, error) {
	var wallet = Wallet{ID: walletID}

	err := db.First(&wallet).Error

	if err != nil {
		return nil, err
	}

	return &wallet.Balance, nil
}

func AddCredit(walletID int, amount int) (*Wallet, error) {
	var wallet = Wallet{ID: walletID}

	err := db.First(&wallet).Error

	if err != nil {
		return nil, err
	}

	currentBalance, _ := strconv.Atoi(wallet.Balance)

	wallet.Balance = strconv.Itoa(currentBalance + amount)

	db.Save(&wallet)

	return &wallet, nil
}

func AddDebit(walletID int, amount int) (*Wallet, error) {
	var wallet = Wallet{ID: walletID}

	err := db.First(&wallet).Error

	if err != nil {
		return nil, err
	}

	currentBalance, _ := strconv.Atoi(wallet.Balance)

	if currentBalance-amount < 0 {
		return nil, errors.New("the Balance is not enough")
	}

	wallet.Balance = strconv.Itoa(currentBalance - amount)

	db.Save(&wallet)

	return &wallet, nil
}
