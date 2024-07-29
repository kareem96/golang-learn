package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)


func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "karim").Info("Hello World")
	logger.WithField("username", "abdul").WithField("name", "Abdul Karim").Info("Hello World")
}
func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "kareem",
		"name": "Abdul karim",
	}).Info("Hello World")
}