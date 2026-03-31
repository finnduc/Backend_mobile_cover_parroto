package logger

import (
	"go-familytree/pkg/settings"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger creates a Zap logger with lumberjack file rotation.

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(cfg settings.LoggerSetting) *zap.Logger {
	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	default:
		level = zapcore.InfoLevel
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,    // MB
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,     // days
		Compress:   cfg.Compress,
	})
	consoleWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		// Always persist every log level to file for troubleshooting/audit.
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), fileWriter, zapcore.DebugLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), consoleWriter, level),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
}
