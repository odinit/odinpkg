package glog

import (
	"github.com/natefinch/lumberjack"
	"github.com/odinit/global/gvar"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

// Init 初始化Logger
func Init(mode, infoPath, errPath string) (err error) {
	//定义日志级别
	//将等级大于等于infoLevel的日志视为一个级别（zap默认有7个日志级别）
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	var core zapcore.Core
	switch mode {
	case "dev":
		// 进入开发模式，日志输出到终端
		config := zap.NewDevelopmentEncoderConfig()
		config.EncodeTime = zapcore.TimeEncoderOfLayout(gvar.DateFormat)
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	case "prod":
		core = zapcore.NewTee(
			//每个zap core 需要三个参数
			//1 Encoder:编码器(如何写入日志)
			//2 WriterSyncer ：指定日志将写到哪里去
			//3 Log Level：哪种级别的日志将被写入
			//zapcore.NewCore(getEncoder(), getLogWriter("./logs/debug.log"), debugLevel),
			zapcore.NewCore(getEncoder(), getLogWriter(infoPath), infoLevel),
			zapcore.NewCore(getEncoder(), getLogWriter(errPath), errorLevel),
		)
	default:
		// 进入开发模式，日志输出到终端
		config := zap.NewDevelopmentEncoderConfig()
		config.EncodeTime = zapcore.TimeEncoderOfLayout(gvar.DateFormat)
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	}

	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger)
	zap.L().Info("日志程序初始化完成")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(gvar.DateFormat)
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, //Filename: 日志文件的位置
		MaxSize:    10,       //MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 10,       //MaxBackups：保留旧文件的最大个数
		MaxAge:     30,       //MaxAges：保留旧文件的最大天数
		Compress:   false,    //Compress：是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
