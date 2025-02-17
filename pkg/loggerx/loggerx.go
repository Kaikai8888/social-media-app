package loggerx

import (
	"context"
	"social-media-app/pkg/trace"

	"go.uber.org/zap"
)

const (
	FieldTraceId = trace.FieldTraceId
)

type Logger interface {
	Debug(ctx context.Context, msg string, fields ...Field)
	Info(ctx context.Context, msg string, fields ...Field)
	Warn(ctx context.Context, msg string, fields ...Field)
	Error(ctx context.Context, msg string, fields ...Field)
	Fatal(ctx context.Context, msg string, fields ...Field)
}

type zapLogger struct {
	l *zap.Logger
}

func NewZapLogger() Logger {
	return &zapLogger{
		l: zap.NewExample(),
	}
}

func (z *zapLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, String(FieldTraceId, trace.GetTraceId(ctx)))
	z.l.Debug(msg, toZapFields(fields)...)
}

func (z *zapLogger) Info(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, String(FieldTraceId, trace.GetTraceId(ctx)))
	z.l.Info(msg, toZapFields(fields)...)
}

func (z *zapLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, String(FieldTraceId, trace.GetTraceId(ctx)))
	z.l.Warn(msg, toZapFields(fields)...)
}

func (z *zapLogger) Error(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, String(FieldTraceId, trace.GetTraceId(ctx)))
	z.l.Error(msg, toZapFields(fields)...)
}

func (z *zapLogger) Fatal(ctx context.Context, msg string, fields ...Field) {
	fields = append(fields, String(FieldTraceId, trace.GetTraceId(ctx)))
	z.l.Fatal(msg, toZapFields(fields)...)
}

func toZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Field(f)
	}
	return zapFields
}
