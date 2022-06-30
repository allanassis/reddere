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
	Bool   = zap.Bool
	String = zap.String
	Int    = zap.Int
	Any    = zap.Any
)

type Logger struct {
	l     *zap.Logger
	level Level
}

func NewLogger(config *config.Config) *Logger {
	level := getLevel(config)
	writer := getWriter()

	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		zapcore.Level(level),
	)
	logger := &Logger{
		l:     zap.New(core),
		level: level,
	}
	return logger
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
