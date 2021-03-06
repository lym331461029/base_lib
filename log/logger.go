package log

import (
	"errors"
	"fmt"
	"github.com/lym331461029/base_lib/log/zap"
	"time"

	"github.com/lym331461029/base_lib/log/base"
	"github.com/lym331461029/base_lib/log/logrus"
)

// DLogger 会返回一个新的默认日志记录器。
func DLogger() base.MyLogger {
	return Logger(base.LOGRUS)
}

// Logger 会新建一个日志记录器并返回。
func Logger(loggerType base.LoggerType) base.MyLogger {
	return LoggerByProjectName(loggerType, "")
}

// LoggerByProjectName 会新建一个日志记录器并返回，同时会接受一个项目名参数并在日志记录中体现。
func LoggerByProjectName(loggerType base.LoggerType, projectName string) base.MyLogger {
	var logger base.MyLogger
	switch loggerType {
	case base.LOGRUS:
		logger = logrus.NewLogger(projectName)
	case base.ZAP:
		logger = zap.NewLogger(projectName)
	default:
		errMsg := fmt.Sprintf("Unsupported logger type '%s'!\n", loggerType)
		panic(errors.New(errMsg))
	}
	if logger != nil {
		base.CheckDebugEnable(projectName, time.Second*10, false)
	}
	return logger
}
