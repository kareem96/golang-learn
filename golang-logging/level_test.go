package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T)  {
	logger := logrus.New()

	logger.Trace("Hello Logger")
	logger.Debug("Hello Logger")
	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}

func TestLoggingLevel(t *testing.T)  {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Hello Logger")
	logger.Debug("Hello Logger")
	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}