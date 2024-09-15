package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() {
    // create new instance 
	log = logrus.New()
	// set lebel to info
	log.SetLevel(logrus.InfoLevel)
    logsDir := "logs/"
    if err := os.MkdirAll(logsDir, 0755); err != nil {
        log.Fatal(err)
    }

    // generate filename with current date and time
    now := time.Now()
    filename := fmt.Sprintf("%s/%d-%02d-%02d_%02d-%02d-%02d_wallet_log.log",
        logsDir,
        now.Year(), now.Month(), now.Day(),
        now.Hour(), now.Minute(), now.Second())
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }

    // config logger
    log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })

    // combine outputs to cli and log files
    log.SetOutput(io.MultiWriter(file, os.Stdout))
}

func GetLogger() *logrus.Logger {
    if log == nil {
        InitLogger()
    }
    return log
}
