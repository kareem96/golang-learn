package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)


func TestSingleton(t *testing.T) {
	logrus.Info("Hello Info")
	logrus.Info("Hello Info")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello Info")
	logrus.Warn("Hello Warn")
}