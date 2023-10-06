package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.ConsoleSeparator = "\n"
	globalsLogger, _ := config.Build()
	zap.ReplaceGlobals(globalsLogger) // 替换zap.s() 作为全局的日志组件
}
