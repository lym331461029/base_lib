package zap // import "go.pkg.wesai.com/p/base_lib/log/zap"


/*
import (
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/lym331461029/base_lib/log/base"
	"github.com/lym331461029/base_lib/log/base/field"
)

func init() {
}

type logger_zap struct {
	projectName string
	inner       zap.Logger
}

// 新建并返回一个日志记录器。
func NewLogger(projectName string) base.MyLogger {
	return NewLoggerBy(projectName, os.Stderr, zap.Debug)
}

// 根据指定的参数新建并返回一个日志记录器。
func NewLoggerBy(projectName string, w io.Writer, l zap.Level) base.MyLogger {
	return &logger_zap{
		projectName: projectName,
		inner:       initInnerLogger(w, l),
	}
}

func initInnerLogger(w io.Writer, l zap.Level) zap.Logger {
	writeSyncer := zap.AddSync(w)
	innerLogger := zap.NewJSON(
		zap.Output(writeSyncer),
		zap.ErrorOutput(writeSyncer),
	)
	innerLogger.SetLevel(l)
	return innerLogger
}

func (logger *logger_zap) Name() string {
	return "zap"
}

func (logger *logger_zap) Debug(v ...interface{}) {
	if base.DebugEnable(logger.projectName) {
		appendRequiredFields(logger.inner).Debug(genMsg("", v...))
	}
}

func (logger *logger_zap) Debugf(format string, v ...interface{}) {
	if base.DebugEnable(logger.projectName) {
		appendRequiredFields(logger.inner).Debug(genMsg(format, v...))
	}
}

func (logger *logger_zap) Debugln(v ...interface{}) {
	if base.DebugEnable(logger.projectName) {
		appendRequiredFields(logger.inner).Debug(genMsg("", v...))
	}
}

func (logger *logger_zap) Error(v ...interface{}) {
	appendRequiredFields(logger.inner).Error(genMsg("", v...))
}

func (logger *logger_zap) Errorf(format string, v ...interface{}) {
	appendRequiredFields(logger.inner).Error(genMsg(format, v...))
}

func (logger *logger_zap) Errorln(v ...interface{}) {
	appendRequiredFields(logger.inner).Error(genMsg("", v...))
}

func (logger *logger_zap) Fatal(v ...interface{}) {
	appendRequiredFields(logger.inner).Fatal(genMsg("", v...))
}

func (logger *logger_zap) Fatalf(format string, v ...interface{}) {
	appendRequiredFields(logger.inner).Fatal(genMsg(format, v...))
}

func (logger *logger_zap) Fatalln(v ...interface{}) {
	appendRequiredFields(logger.inner).Fatal(genMsg("", v...))
}

func (logger *logger_zap) Info(v ...interface{}) {
	appendRequiredFields(logger.inner).Info(genMsg("", v...))
}

func (logger *logger_zap) Infof(format string, v ...interface{}) {
	appendRequiredFields(logger.inner).Info(genMsg(format, v...))
}

func (logger *logger_zap) Infoln(v ...interface{}) {
	appendRequiredFields(logger.inner).Info(genMsg("", v...))
}

func (logger *logger_zap) Panic(v ...interface{}) {
	appendRequiredFields(logger.inner).Panic(genMsg("", v...))
}

func (logger *logger_zap) Panicf(format string, v ...interface{}) {
	appendRequiredFields(logger.inner).Panic(genMsg(format, v...))
}

func (logger *logger_zap) Panicln(v ...interface{}) {
	appendRequiredFields(logger.inner).Panic(genMsg("", v...))
}

func (logger *logger_zap) Warn(v ...interface{}) {
	appendRequiredFields(logger.inner).Warn(genMsg("", v...))
}

func (logger *logger_zap) Warnf(format string, v ...interface{}) {
	appendRequiredFields(logger.inner).Warn(genMsg(format, v...))
}

func (logger *logger_zap) Warnln(v ...interface{}) {
	appendRequiredFields(logger.inner).Warn(genMsg("", v...))
}

func (logger *logger_zap) WithFields(fields ...field.Field) base.MyLogger {
	fieldsLen := len(fields)
	if fieldsLen == 0 {
		return logger
	}
	zapFields := make([]zap.Field, 0)
	for _, curfield := range fields {
		switch curfield.Type() {
		case field.BoolType:
			zapFields = append(zapFields,
				zap.Bool(curfield.Name(), false))
		case field.Int64Type:
			zapFields = append(zapFields,
				zap.Int64(curfield.Name(), curfield.Value().(int64)))
		case field.Float64Type:
			zapFields = append(zapFields,
				zap.Float64(curfield.Name(), curfield.Value().(float64)))
		case field.StringType:
			zapFields = append(zapFields,
				zap.String(curfield.Name(), curfield.Value().(string)))
		case field.ObjectType:
			zapFields = append(zapFields,
				zap.Object(curfield.Name(), curfield.Value().(string)))
		default:
			continue
		}
	}
	return &logger_zap{
		projectName: logger.projectName,
		inner:       logger.inner.With(zapFields...),
	}
}

// 添加必要的字段。
func appendRequiredFields(logger zap.Logger) zap.Logger {
	return appendLocation(appendTimeString(logger))
}

// 添加时间的字符串形式。
func appendTimeString(logger zap.Logger) zap.Logger {
	timeString := time.Now().Format(base.TIMESTAMP_FORMAT)
	field := zap.String("time", timeString)
	return logger.With(field)
}

// 添加记录日志的代码位置。
func appendLocation(logger zap.Logger) zap.Logger {
	funcPath, fileName, line := base.GetInvokerLocation(4)
	field := zap.Nest(
		"location",
		zap.String("func_path", funcPath),
		zap.String("file_name", fileName),
		zap.Int("line", line),
	)
	return logger.With(field)
}

// 生成日志消息。
func genMsg(format string, v ...interface{}) string {
	if len(format) > 0 {
		return fmt.Sprintf(format, v...)
	} else {
		return fmt.Sprint(v...)
	}
}
*/