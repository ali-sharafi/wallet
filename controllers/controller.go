package controllers

import (
	"fmt"

	"github.com/ali-sharafi/wallet/models"
	"github.com/gin-gonic/gin"
)

func GetWallets(c *gin.Context) {
	wallets, err := models.GetWallets()
	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, wallets)
}

func GetBalance(c *gin.Context) {
	fmt.Println("Here is GetBalance")
}

func AddCredit(c *gin.Context) {
	fmt.Println("Here is AddCredit")
}

func AddDebit(c *gin.Context) {
	fmt.Println("Here is AddDebit")
}
