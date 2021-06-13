package internal

import (
	"io"
	"os"

	"easygo/internal/common/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getLogWriter() io.Writer {
	c := config.C.LogFileHook

	hook := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}

	return hook
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// InitLogger 日志初始化
func InitLogger() func() {
	c := config.C.Log

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(c.Level))

	var writer zapcore.WriteSyncer
	if c.Output == "file" {
		writer = zapcore.AddSync(getLogWriter())
	} else {
		writer = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(
		getEncoder(),
		writer,
		atomicLevel,
	)

	var tempLogger *zap.Logger
	if config.C.Mode == "dev" {
		tempLogger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		tempLogger = zap.New(core)
	}

	undo := zap.ReplaceGlobals(tempLogger)

	return func() {
		undo()
	}
}
