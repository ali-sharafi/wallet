package routes

import (
	"fmt"

	"github.com/ali-sharafi/wallet/controllers"
	"github.com/ali-sharafi/wallet/pkg/logging"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(RequestLogger())

	apiV1 := r.Group("/api/v1/wallets")

	apiV1.GET("/", controllers.GetWallets)

	apiV1.GET("/:id/balance", controllers.GetBalance)

	apiV1.POST("/:id/credit", controllers.AddCredit)

	apiV1.POST("/:id/debit", controllers.AddDebit)

	return r
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logging.Log(fmt.Sprintf("New request from %v", c.Request.URL))
		c.Next()
	}
}
