package logrus

import (
	"os"
	"time"

	"github.com/Becklyn/go-fx-core/env"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx/fxevent"
)

var LOG_LEVEL = "LOG_LEVEL"

func newLogrus(_ *env.Env) *logrus.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC822,
	})
	logger.SetLevel(getLogLevel())
	logger.WithFields(logrus.Fields{
		"environment": env.String(env.APP_ENV),
	}).Info("Using environment")

	return logger
}

func logrusFxLogger(logger *logrus.Logger) fxevent.Logger {
	if !env.IsDevelopment() {
		return fxevent.NopLogger
	}
	return &fxevent.ConsoleLogger{W: logger.Writer()}
}

func getLogLevel() logrus.Level {
	switch env.String(LOG_LEVEL) {
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
