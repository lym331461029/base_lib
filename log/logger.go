package log

import (
	"errors"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lym331461029/base_lib/log/zap"
	"os"
	"path"
	"time"

	"github.com/lym331461029/base_lib/log/base"
	"github.com/lym331461029/base_lib/log/logrus"
	rLogrus "github.com/sirupsen/logrus"
	rZapcore "go.uber.org/zap/zapcore"
)

// DLogger 会返回一个新的默认日志记录器。
func DLogger() base.MyLogger {
	return Logger(base.LOGRUS)
}

// Logger 会新建一个日志记录器并返回。
func Logger(loggerType base.LoggerType) base.MyLogger {
	return LoggerByProjectName(loggerType, "")
}

// LoggerToFile 会新建一个日志记录器并返回。
func LoggerToFile(loggerType base.LoggerType, projectName string, dirPath string) base.MyLogger {
	l, err := loggerToFile(loggerType, projectName, dirPath)
	if err != nil {
		panic(err.Error())
	}
	return l
}

func loggerToFile(loggerType base.LoggerType, projectName string, dirPath string) (base.MyLogger, error) {
	if len(dirPath) <= 0 {
		dirPath, _ = os.Getwd()
	}
	_, err := os.Stat(dirPath)
	if err != nil {
		if err == os.ErrNotExist {
			err = os.Mkdir(dirPath, os.ModePerm)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	filePath := path.Join(dirPath, projectName+".%Y%m%d%H.log")

	wlogs, err := rotatelogs.New(filePath, rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		return nil, err
	}

	var logger base.MyLogger
	switch loggerType {
	case base.LOGRUS:
		logger = logrus.NewLoggerBy(projectName, wlogs, rLogrus.DebugLevel)
	case base.ZAP:
		logger = zap.NewLoggerBy(projectName, wlogs, rZapcore.DebugLevel)
	default:
		errMsg := fmt.Sprintf("Unsupported logger type '%s'!\n", loggerType)
		panic(errors.New(errMsg))
	}
	return logger, nil
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
