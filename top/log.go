package top

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

var core = zapcore.NewNopCore()
var Logger *zap.Logger
var SLogger *zap.SugaredLogger

func NewLogger(logMode, logFilePath string) *zap.Logger {
	if slices.Contains([]string{"release", "product", "prod", "pro", "production"}, strings.ToLower(logMode)) {
		return product(logFilePath)
	} else {
		return develop()
	}
}

func GlobalLog(logMode, logFilePath string) {
	Logger = NewLogger(logMode, logFilePath)
	SLogger = Logger.Sugar()
}

// develop 开发者模式
// 日志仅输出到终端
func develop() *zap.Logger {
	devConfig := zap.NewDevelopmentEncoderConfig()
	devConfig.EncodeTime = zapcore.TimeEncoderOfLayout(DateFormat)

	consoleEncoder := zapcore.NewConsoleEncoder(devConfig)
	core = zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)

	return zap.New(core, zap.AddCaller())
}

// product 生产者模式
// 日志输出到文件
func product(logFile string) *zap.Logger {
	proConfig := zap.NewProductionEncoderConfig()
	proConfig.EncodeTime = zapcore.TimeEncoderOfLayout(DateFormat)
	proConfig.TimeKey = "time"
	proConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	proConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	proConfig.EncodeCaller = zapcore.ShortCallerEncoder

	consoleEncoder := zapcore.NewJSONEncoder(proConfig)
	logWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile, //Filename: 日志文件的位置
		MaxSize:    10,      //MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,      //MaxBackups：保留旧文件的最大个数
		MaxAge:     30,      //MaxAges：保留旧文件的最大天数
		Compress:   false,   //Compress：是否压缩/归档旧文件
	})

	core = zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, logWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel
		})),
	)

	return zap.New(core, zap.AddCaller())
}
