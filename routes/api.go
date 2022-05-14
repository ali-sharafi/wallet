package routes

import (
	"github.com/ali-sharafi/wallet/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiV1 := r.Group("/api/v1/wallets")

	apiV1.GET("/", controllers.GetWallets)

	apiV1.GET("/:id/balance", controllers.GetBalance)

	apiV1.POST("/:id/credit", controllers.AddCredit)

	apiV1.POST("/:id/debit", controllers.AddDebit)

	return r
}
