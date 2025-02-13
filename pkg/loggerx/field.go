package loggerx

import "go.uber.org/zap"

type Field zap.Field

func String(key string, val string) Field {
	return Field(zap.String(key, val))
}

func Int(key string, val int) Field {
	return Field(zap.Int(key, val))
}

func Int64(key string, val int64) Field {
	return Field(zap.Int64(key, val))
}

func Float64(key string, val float64) Field {
	return Field(zap.Float64(key, val))
}

func Bool(key string, val bool) Field {
	return Field(zap.Bool(key, val))
}

func Any(key string, val interface{}) Field {
	return Field(zap.Any(key, val))
}
