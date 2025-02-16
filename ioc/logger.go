package ioc

import "webook/pkg/loggerx"

func InitLogger() loggerx.Logger {
	return loggerx.NewZapLogger()
}
