package utils

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, message string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Msg:  message,
		Data: data,
	})
	return
}
