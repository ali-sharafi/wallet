package controllers

import (
	"net/http"
	"strconv"

	"github.com/ali-sharafi/wallet/models"
	"github.com/ali-sharafi/wallet/utils"
	"github.com/gin-gonic/gin"
)

var form struct {
	Amount   int `json:"amount"`
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

	balance, err := models.GetBalance(walletID)

	if err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

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

	if form.Amount < 1 || err != nil {
		return false, "the amount value must be greater than zero"
	}

	return true, ""
}
