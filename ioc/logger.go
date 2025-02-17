package ioc

import "social-media-app/pkg/loggerx"

func InitLogger() loggerx.Logger {
	return loggerx.NewZapLogger()
}
