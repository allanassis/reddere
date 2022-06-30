package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	InfoLevel  Level = zap.InfoLevel  // 0, default level
	WarnLevel  Level = zap.WarnLevel  // 1
	ErrorLevel Level = zap.ErrorLevel // 2
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)
