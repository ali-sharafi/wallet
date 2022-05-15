package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ali-sharafi/wallet/pkg/file"
)

var (
	logger *log.Logger
	F      *os.File
)

func Setup() {
	var err error
	fileName := getLogFileName()
	filePath := getLogFilePath()
	F, err = file.MustOpen(fileName, filePath)

	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, "", log.LstdFlags)
}

func Log(v ...interface{}) {
	logger.Println(v)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", "runtime/", "logs/")
}

func getLogFileName() string {
	return fmt.Sprintf("%s.%s",
		time.Now().Format("2006-01-02"),
		"txt",
	)
}
