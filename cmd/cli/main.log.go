package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	// customize
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Info("Info log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// trans 1733415346.2475882 --> 2024-12-05T23:15:46.247+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfig.TimeKey = "ts"
	//  from info --> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller": "cli/log/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	// Ensure the directory exists
	if err := os.MkdirAll("./log", os.ModePerm); err != nil {
		panic(err)
	}

	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
