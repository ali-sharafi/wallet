package routes

import (
	"fmt"
	"net/http"

	"github.com/ali-sharafi/wallet/controllers"
	"github.com/ali-sharafi/wallet/pkg/logging"
	"github.com/ali-sharafi/wallet/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(RequestLogger())

	r.POST("/auth", controllers.GetAuth)

	apiV1 := r.Group("/api/v1/wallets")
	apiV1.Use(JWT())

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

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var err error
		var claims *jwt.StandardClaims

		code = http.StatusOK
		token := c.GetHeader("Authorization")

		if token == "" {
			code = http.StatusUnauthorized
		} else {
			claims, err = utils.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "unauthorized",
				"data": "",
			})

			c.Abort()
			return
		}

		c.Set("user_id", claims.Issuer)
		c.Next()
	}
}
