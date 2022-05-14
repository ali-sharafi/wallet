package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ali-sharafi/wallet/models"
	"github.com/ali-sharafi/wallet/utils"
	"github.com/gin-gonic/gin"
)

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
	var form struct {
		Amount int `json:"amount"`
	}

	walletID, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		res.Response(http.StatusBadRequest, "Invalid Params", nil)
		return
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	if form.Amount < 1 || err != nil {
		res.Response(http.StatusBadRequest, "Amount is not valid", nil)
		return
	}

	wallet, err := models.AddCredit(walletID, form.Amount)

	if err != nil {
		res.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}

	res.Response(http.StatusOK, "Success", wallet)
}

func AddDebit(c *gin.Context) {
	fmt.Println("Here is AddDebit")
}
