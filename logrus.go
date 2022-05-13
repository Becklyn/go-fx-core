package core

import (
	"os"
	"time"

	"github.com/Becklyn/go-fx-core/env"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

var LOG_LEVEL = "LOG_LEVEL"

var LogrusModule = fx.Options(
	fx.Provide(
		newLogrus,
	),
	fx.WithLogger(logrusFxLogger),
)

func newLogrus() *logrus.Logger {
	logger := logrus.New()

	logrus.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC822,
	})
	logger.SetLevel(getLogLevel())
	logger.Infof("Using %s environment", env.String(env.APP_ENV))

	return logger
}

func logrusFxLogger(logger *logrus.Logger) fxevent.Logger {
	if env.IsDevelopment() {

	}

	return &fxevent.ConsoleLogger{W: logger.Writer()}
}

func getLogLevel() logrus.Level {
	switch env.String(LOG_LEVEL) {
	case "info":
		return logrus.InfoLevel
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