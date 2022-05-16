package models

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID        int             `gorm:"primary_key" json:"id"`
	Balance   decimal.Decimal `json:"balance"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
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

func GetBalance(walletID int) (*decimal.Decimal, error) {
	var wallet = Wallet{ID: walletID}

	err := db.First(&wallet).Error

	if err != nil {
		return nil, err
	}

	return &wallet.Balance, nil
}

func AddCredit(walletID int, amount string) (*Wallet, error) {
	var wallet = Wallet{ID: walletID}

	err := db.First(&wallet).Error

	if err != nil {
		return nil, err
	}

	newAmount, err := decimal.NewFromString(amount)

	if err != nil {
		return nil, err
	}

	wallet.Balance = wallet.Balance.Add(newAmount)

	db.Save(&wallet)

	return &wallet, nil
}

func AddDebit(walletID int, amount string) (*Wallet, error) {
	var wallet = Wallet{ID: walletID}

	err := db.First(&wallet).Error

	if err != nil {
		return nil, err
	}

	newAmount, err := decimal.NewFromString(amount)

	if err != nil {
		return nil, err
	}

	if wallet.Balance.Sub(newAmount).IsNegative() {
		return nil, errors.New("the Balance is not enough")
	}

	wallet.Balance = wallet.Balance.Sub(newAmount)

	db.Save(&wallet)

	return &wallet, nil
}
