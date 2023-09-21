package common

// 重新构建logger

import (
	"io"
	"os"
	"time"

	"github.com/shinesyang/go_logger/logrotate"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

var infoFilename = logrotate.DefaultFileName
var warnFilename = "error.log"

type LoggerParam struct {
	stdout   io.Writer
	stderr   io.Writer
	infoFile string
	warnFile string
}

func init() {
	InitLogger()
}

func NewCore(stdout, stderr io.Writer) zapcore.Core {
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	/*Info级别以下,Debug 不输出到控制台*/
	infoEqualLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(getEncoder(), getWriteInfoSyncer(), infoLevel),
		zapcore.NewCore(getEncoder(), getWriteWarnSyncer(), warnLevel),
		zapcore.NewCore(getEncoder(), zapcore.AddSync(stdout), infoEqualLevel),    // 日志生成到标注输出(控制台)
		zapcore.NewCore(getEncoder(), zapcore.AddSync(stderr), zapcore.WarnLevel), // 警告级别的日志生成带标准错误(控制台)
	)
	return core

}

func InitLogger() {
	core := NewCore(os.Stdout, os.Stderr)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
}

func CustomLogger(param LoggerParam) *zap.SugaredLogger {
	if param.infoFile != "" {
		infoFilename = param.infoFile
	}

	if param.warnFile != "" {
		warnFilename = param.warnFile
	}

	var stdout io.Writer
	var stderr io.Writer
	if param.stdout != nil {
		stdout = param.stdout
	} else {
		stdout = os.Stdout
	}

	if param.stderr != nil {
		stderr = param.stderr
	} else {
		stderr = os.Stderr
	}

	core := NewCore(stdout, stderr)
	return zap.New(core, zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Local().Format("2006-01-02 15:04:05"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriteInfoSyncer() zapcore.WriteSyncer {
	r, _ := logrotate.NewRotateLog(
		logrotate.WithRotateFilePath(logrotate.DefaultFilePath, infoFilename),
		logrotate.WithDeleteExpiredFile(logrotate.MaxAgeQuarter),
		logrotate.WithSimpleControl(),
	)
	return zapcore.AddSync(r)
}

func getWriteWarnSyncer() zapcore.WriteSyncer {
	r, _ := logrotate.NewRotateLog(
		logrotate.WithRotateFilePath(logrotate.DefaultFilePath, warnFilename),
		logrotate.WithDeleteExpiredFile(logrotate.MaxAgeQuarter),
		logrotate.WithSimpleControl(),
	)
	return zapcore.AddSync(r)
}

func getLevelEnabler() zapcore.LevelEnabler {

	//return zapcore.DebugLevel
	return zapcore.InfoLevel
}
