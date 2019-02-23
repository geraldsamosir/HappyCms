package Config

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	LOGPATH = "server/logs"
)

type Log struct {
}

func (l *Log) Logging() {
	file, err := os.OpenFile(LOGPATH+"/info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
