package global

import (
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var mainLog *zap.Logger = nil

func Print(s string) {
	mainLog.Debug(s)
}

func Error(s string) {
	mainLog.Error(s)
}

func Fatal(s string) {
	mainLog.Fatal(s)
}

func init() {
	mainLog = initLogSystem()
}

func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func initLogSystem() *zap.Logger {
	core := newLogCore("logfile.json", zapcore.InfoLevel, 128, 30, 7, true)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("ServerName", "KormanServer")))
}

func newLogCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	hook := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktracer",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	return zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync((&hook))), atomicLevel)
}
