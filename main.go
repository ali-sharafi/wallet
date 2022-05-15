package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ali-sharafi/wallet/models"
	"github.com/ali-sharafi/wallet/pkg/logging"
	"github.com/ali-sharafi/wallet/pkg/settings"
	"github.com/ali-sharafi/wallet/routes"
)

func init() {
	settings.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	routeHandler := routes.InitRouter()

	endPoint := fmt.Sprintf("localhost:%d", settings.ServerSetting.HttpPort)
	readTimeout := settings.ServerSetting.ReadTimeout
	writeTimeout := settings.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Handler:        routeHandler,
		Addr:           endPoint,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
