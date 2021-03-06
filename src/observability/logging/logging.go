package logging

import (
	"io"
	"os"

	"github.com/allanassis/reddere/src/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

var (
	Bool     = zap.Bool
	String   = zap.String
	Duration = zap.Duration
	Int      = zap.Int
	Int64    = zap.Int64
	Any      = zap.Any
)

type Logger struct {
	l     *zap.Logger
	level Level
}

func NewLogger(config *config.Config) *Logger {
	level := getLevel(config)
	writer := getWriter()

	zapLogger := newZapLogger(level, writer)

	logger := &Logger{
		l:     zapLogger,
		level: level,
	}
	return logger
}

func newZapLogger(level Level, writer io.Writer) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)

	return zap.New(core)
}

func (l *Logger) With(fields ...Field) *Logger {
	newLogger := l.l.With(fields...)
	return &Logger{
		l:     newLogger,
		level: l.level,
	}
}

func (l *Logger) UpdateFiels(fields ...Field) {
	l.l = l.l.With(fields...)
}

func (l *Logger) CleanUpFields() {
	writer := getWriter()
	zapLogger := newZapLogger(l.level, writer)

	l.l = zapLogger
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func getLevel(config *config.Config) Level {
	level := config.GetString("observability.logging.level")
	levelsMap := map[string]Level{
		"info":  InfoLevel,
		"debug": DebugLevel,
		"warn":  WarnLevel,
		"error": ErrorLevel,
		"fatal": FatalLevel,
	}
	return levelsMap[level]
}

func getWriter() io.Writer {
	return os.Stdout
}
