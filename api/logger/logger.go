package logger

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var loggingEnabled bool

func Init(enabled bool) {
	loggingEnabled = enabled
}

func Log(fields log.Fields, message string) {
	if loggingEnabled {
		customFormatter := new(logrus.TextFormatter)
		customFormatter.TimestampFormat = "2006-01-02 15:04:05"
		customFormatter.FullTimestamp = true
		customFormatter.PadLevelText = true
		log.SetFormatter(customFormatter)
		log.WithFields(fields).Error(message)
	}
}
