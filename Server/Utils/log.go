package Utils

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	LOGPATH = "server/logs"
)

type Log struct {
}

func (l *Log) Logging(message interface{}, messageErr string) {
	file, err := os.OpenFile(LOGPATH+"/info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	errResp, _ := json.Marshal(message)
	logrus.WithFields(logrus.Fields{
		"responsetoFrontend": errResp,
	}).Info(messageErr)
	file.Close()
}
