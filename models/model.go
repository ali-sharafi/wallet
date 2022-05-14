package models

import (
	"fmt"
	"log"
	"time"

	"github.com/ali-sharafi/wallet/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(settings.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DatabaseSetting.User,
		settings.DatabaseSetting.Password,
		settings.DatabaseSetting.Host,
		settings.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("setupDB err: %v", err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	migrate()
	seedTable()
}

func seedTable() {
	dateTimeFormat := "2006-01-02 15:04:05"
	wallets := []Wallet{
		{ID: 1, Balance: "1000", CreatedAt: time.Now().Format(dateTimeFormat), UpdatedAt: time.Now().Format(dateTimeFormat)},
		{ID: 2, Balance: "1000", CreatedAt: time.Now().Format(dateTimeFormat), UpdatedAt: time.Now().Format(dateTimeFormat)},
		{ID: 3, Balance: "1000", CreatedAt: time.Now().Format(dateTimeFormat), UpdatedAt: time.Now().Format(dateTimeFormat)},
	}

	db.Delete(&Wallet{})
	for _, wallet := range wallets {
		db.Create(&Wallet{ID: wallet.ID, Balance: wallet.Balance, CreatedAt: wallet.CreatedAt, UpdatedAt: wallet.UpdatedAt})
	}
}

func migrate() {
	db.AutoMigrate(&Wallet{})
}
