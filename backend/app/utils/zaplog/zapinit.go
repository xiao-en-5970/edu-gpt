package zaplog

import (
	"os"
	"github.com/xiao-en-5970/edu-gpt/backend/app/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2" // 引入lumberjack库
)

func InitZapLogger() {
	// 1. 配置日志文件滚动（使用lumberjack）
	logFile := &lumberjack.Logger{
		Filename:   global.Cfg.Logging.FilePath+"/log.log", // 日志文件路径
		MaxSize:    global.Cfg.Logging.MaxSize,  // 单个日志文件最大大小（MB）
		MaxBackups: global.Cfg.Logging.MaxBackups, // 保留的旧日志文件数量
		MaxAge:     global.Cfg.Logging.MaxAge,     // 保留旧日志文件的最大天数
		Compress:   global.Cfg.Logging.Compress,   // 是否压缩旧日志文件
	}

	// 2. 设置日志级别（文件输出）
	var fileLevel zapcore.Level
	switch global.Cfg.Logging.FileLevel {
	case "debug":
		fileLevel = zap.DebugLevel
	case "info":
		fileLevel = zap.InfoLevel
	case "warn":
		fileLevel = zap.WarnLevel
	case "error":
		fileLevel = zap.ErrorLevel
	case "fatal":
		fileLevel = zap.FatalLevel
	default:
		fileLevel = zap.WarnLevel
	}

	// 3. 文件编码器配置（JSON格式，无颜色）
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.TimeKey = "timestamp"
	fileEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)

	// 4. 文件输出核心
	fileCore := zapcore.NewCore(
		fileEncoder,
		zapcore.AddSync(logFile), // 使用lumberjack实现文件滚动
		fileLevel,
	)

	// 5. 设置日志级别（控制台输出）
	var consoleLevel zapcore.Level
	switch global.Cfg.Logging.StdOutLevel {
	case "debug":
		consoleLevel = zap.DebugLevel
	case "info":
		consoleLevel = zap.InfoLevel
	case "warn":
		consoleLevel = zap.WarnLevel
	case "error":
		consoleLevel = zap.ErrorLevel
	case "fatal":
		consoleLevel = zap.FatalLevel
	default:
		consoleLevel = zap.WarnLevel
	}

	// 6. 控制台编码器配置（带颜色）
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig() // 使用开发模式配置（默认带颜色）
	consoleEncoderConfig.TimeKey = "timestamp"
	consoleEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 彩色日志级别
	consoleEncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder    // 简短调用者信息

	// 关键修正：使用NewConsoleEncoder创建控制台编码器
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig) // ✅ 正确：使用ConsoleEncoder

	// 7. 控制台输出核心
	consoleCore := zapcore.NewCore(
		consoleEncoder, // ✅ 传入ConsoleEncoder，而不是JSONEncoder
		zapcore.AddSync(os.Stdout),
		consoleLevel,
	)

	// 8. 组合多个核心（同时输出到文件和控制台）
	core := zapcore.NewTee(fileCore, consoleCore)

	// 9. 创建Logger
	l := zap.New(
		core,
		zap.AddCaller(),           // 记录调用者信息
		zap.AddStacktrace(zapcore.ErrorLevel), // Error级别及以上记录堆栈
	)

	// 10. 替换全局Logger
	global.Logger = l.Sugar()
	global.Logger.Info("Zap logger initialized successfully with log rotation")
}