package logging

import (
	"{{cookiecutter.app_name}}/config"
	"github.com/sirupsen/logrus"
)

// GetLogger :
func GetLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

// Init :
func Init() {
	logger := logrus.StandardLogger()
    level, err := logrus.ParseLevel(config.Configuration.Logging.Level)
    if err != nil {
        panic(err)
	}
	logger.SetLevel(level)
}