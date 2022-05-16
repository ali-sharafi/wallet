package controllers

import (
	"net/http"
	"strconv"

	"github.com/ali-sharafi/wallet/models"
	"github.com/ali-sharafi/wallet/pkg/gredis"
	"github.com/ali-sharafi/wallet/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

var form struct {
	Amount   string `json:"amount"`
	WalletID int
}

func GetWallets(c *gin.Context) {
	res := utils.Gin{C: c}
	wallets, err := models.GetWallets()
	if err != nil {
		res.Response(http.StatusBadRequest, "Something went wrong", nil)
		return
	}

	res.Response(http.StatusOK, "Success", wallets)
}

func GetBalance(c *gin.Context) {
	res := utils.Gin{C: c}
	walletID, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		res.Response(http.StatusBadRequest, "Invalid Wallet", nil)
		return
	}

	balance, err := gredis.Get(strconv.Itoa(walletID))

	if err == nil {
		res.Response(http.StatusOK, "Success", balance)
		return
	}

	balance, err = models.GetBalance(walletID)

	if err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	gredis.Set(strconv.Itoa(walletID), balance, 3600)

	res.Response(http.StatusOK, "Success", balance)
}

func AddCredit(c *gin.Context) {
	res := utils.Gin{C: c}

	result, msg := bindAndValidate(c)

	if !result {
		res.Response(http.StatusBadRequest, msg, nil)
		return
	}

	wallet, err := models.AddCredit(form.WalletID, form.Amount)

	if err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	res.Response(http.StatusOK, "Success", wallet)
}

func AddDebit(c *gin.Context) {
	res := utils.Gin{C: c}

	result, msg := bindAndValidate(c)

	if !result {
		res.Response(http.StatusBadRequest, msg, nil)
		return
	}

	wallet, err := models.AddDebit(form.WalletID, form.Amount)

	if err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	res.Response(http.StatusOK, "Success", wallet)
}

func bindAndValidate(c *gin.Context) (result bool, msg string) {
	walletID, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		return false, "Invalid Params"
	}

	form.WalletID = walletID

	if err := c.ShouldBindJSON(&form); err != nil {
		return false, err.Error()
	}

	if amount, err := decimal.NewFromString(form.Amount); amount.IsNegative() || err != nil {
		return false, "the amount value must be positive"
	}

	return true, ""
}
