package main

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	logger := newLogger()

	logger.Info("Server started", zap.Int("port", 8080))
	logger.Error("Database error", zap.String("db", "postgres"))

}

func newLogger() *zap.Logger {

	encoder := getEncoder()

	write := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/log.log",
		MaxSize:    10,
		MaxAge:     2,
		MaxBackups: 3,
		Compress:   true,
	})

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(write, zapcore.AddSync(os.Stdout)),
		zapcore.InfoLevel,
	)

	return zap.New(core, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {

	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		MessageKey:     "msg",
		CallerKey:      "caller",
		EncodeTime:     customTimeEncoder,
		EncodeLevel:    customLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	return zapcore.NewJSONEncoder(config)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}

func getWriter() zapcore.WriteSyncer {

	os.MkdirAll("./log", os.ModePerm)

	file, _ := os.OpenFile(
		"./log/app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		os.ModePerm,
	)

	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)
}
