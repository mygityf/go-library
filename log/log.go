package log

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// init out logger
func init() {
	SetOutLogger(NewConsoleLoggerImpl())
}

var (
	InStdoutTestMode bool
)

type consoleLoggerImpl struct{}

// NewConsoleLoggerImpl stdout logger
func NewConsoleLoggerImpl() *consoleLoggerImpl {
	return &consoleLoggerImpl{}
}

// Write to stdout
func (c *consoleLoggerImpl) Write(ctx context.Context, level Level, row string) {
	fmt.Println(ctx, level.CapitalString(), row)
}

// Close close
func (c *consoleLoggerImpl) Close() {}

type logChain struct {
	ctx      context.Context
	lines    []string
	funcName string
	message  string
	level    Level
}

// Any log any type message
func (l *logChain) Any(key string, message interface{}) *logChain {
	l.lines = append(l.lines, strings.Join([]string{key, ToString(message)}, ":"))
	return l
}

// Ctx add ctx
func (l *logChain) Ctx(ctx context.Context) *logChain {
	if ctx == nil {
		return l
	}
	l.ctx = ctx
	return l
}

// Error add error
func (l *logChain) Error(err error) *logChain {
	if err == nil {
		return l
	}
	return l.Any("error", err.Error())
}

// Line log out
func (l *logChain) Line() {
	message := "[" + l.funcName + "] " + l.message + "; " + strings.Join(l.lines, ",")
	if InStdoutTestMode {
		fmt.Println(message)
		return
	}
	if defaultOutLogger != nil {
		defaultOutLogger.Write(l.ctx, l.level, message)
	}
}

func ToString(data interface{}) (res string) {
	switch v := data.(type) {
	case bool:
		res = strconv.FormatBool(v)
	case float32:
		res = strconv.FormatFloat(float64(v), 'f', 6, 32)
	case float64:
		res = strconv.FormatFloat(v, 'f', 6, 64)
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(data)
		res = strconv.FormatInt(int64(val.Int()), 10)
	case uint, uint8, uint16, uint32, uint64:
		val := reflect.ValueOf(data)
		res = strconv.FormatUint(uint64(val.Uint()), 10)
	case string:
		res = v
	case []byte:
		res = string(v)
	case error:
		if v != nil {
			res = v.Error()
		}
	default:
		if stringer, ok := data.(fmt.Stringer); ok {
			res = stringer.String()
		} else {
			js, _ := json.Marshal(data)
			res = string(js)
		}
	}
	return
}

// GetFunction get func name for logger
func GetFunction(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	pos := runtime.FuncForPC(pc).Name()
	posList := strings.Split(pos, ".")
	if len(posList) > 0 {
		return posList[len(posList)-1]
	}
	return pos
}

// Debug log.Debug("xxxx).Ctx(ctx).Line()
func Debug(msg string) *logChain {
	return &logChain{
		level:    DebugLevel,
		funcName: GetFunction(2),
		message:  msg,
	}
}

// Info log.Info("xxxx).Ctx(ctx).Line()
func Info(msg string) *logChain {
	return &logChain{
		level:    InfoLevel,
		funcName: GetFunction(2),
		message:  msg,
	}
}

// Warn log.Warn("xxxx).Ctx(ctx).Line()
func Warn(msg string) *logChain {
	return &logChain{
		level:    WarnLevel,
		funcName: GetFunction(2),
		message:  msg,
	}
}

// Err log.Err("xxxx).Ctx(ctx).Error(err).Line()
func Err(msg string) *logChain {
	return &logChain{
		level:    ErrorLevel,
		funcName: GetFunction(2),
		message:  msg,
	}
}

// Fatal log.Fatal("xxxx).Ctx(ctx).Line()
func Fatal(msg string) *logChain {
	return &logChain{
		level:    FatalLevel,
		funcName: GetFunction(2),
		message:  msg,
	}
}

func formatLog(format string, args ...interface{}) string {
	if format != "" {
		return fmt.Sprintf(format, args...)
	} else {
		return fmt.Sprint(args...)
	}
}

// Debugf log.Debugf("xxxx-%v", "x").Ctx(ctx).Line()
func Debugf(fmt string, args ...interface{}) *logChain {
	return &logChain{
		level:    DebugLevel,
		funcName: GetFunction(2),
		message:  formatLog(fmt, args...),
	}
}

// Infof log.Infof("xxxx-%v", "x").Ctx(ctx).Line()
func Infof(fmt string, args ...interface{}) *logChain {
	return &logChain{
		level:    InfoLevel,
		funcName: GetFunction(2),
		message:  formatLog(fmt, args...),
	}
}

// Warnf log.Warnf("xxxx-%v", "x").Ctx(ctx).Line()
func Warnf(fmt string, args ...interface{}) *logChain {
	return &logChain{
		level:    WarnLevel,
		funcName: GetFunction(2),
		message:  formatLog(fmt, args...),
	}
}

// Errorf log.Errorf("xxxx-%v", "x").Ctx(ctx).Error(err).Line()
func Errorf(fmt string, args ...interface{}) *logChain {
	return &logChain{
		level:    ErrorLevel,
		funcName: GetFunction(2),
		message:  formatLog(fmt, args...),
	}
}

// Fatalf log.Fatalf("xxxx-%v", "x").Ctx(ctx).Line()
func Fatalf(fmt string, args ...interface{}) *logChain {
	return &logChain{
		level:    FatalLevel,
		funcName: GetFunction(2),
		message:  formatLog(fmt, args...),
	}
}
