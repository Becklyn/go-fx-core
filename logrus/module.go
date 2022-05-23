package logrus

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		newLogrus,
	),
	fx.WithLogger(logrusFxLogger),
)
