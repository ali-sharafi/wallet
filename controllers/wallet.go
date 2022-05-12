package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetWallets(c *gin.Context) {
	fmt.Println("Here is get wallets")
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
