package logrus

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(1))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
