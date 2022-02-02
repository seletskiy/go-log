package log

import (
	"fmt"
	"regexp"
	"runtime"

	"github.com/kovetskiy/lorg"
	"github.com/reconquest/cog"
	"github.com/reconquest/colorgful"
	"github.com/reconquest/karma-go"
)

type Logger interface {
	SetLevel(Level)
	GetLevel() Level
	NewChild() Logger
	NewChildWithPrefix(string) Logger
	Fatalf(error, string, ...interface{})
	Errorf(error, string, ...interface{})
	Warningf(error, string, ...interface{})
	Infof(*karma.Context, string, ...interface{})
	Debugf(*karma.Context, string, ...interface{})
	Tracef(*karma.Context, string, ...interface{})
	Fatal(values ...interface{})
	Error(values ...interface{})
	Warning(values ...interface{})
	Info(values ...interface{})
	Debug(values ...interface{})
	Trace(values ...interface{})
}

var (
	proxy          Proxy
	implementation *cog.Logger
	theme          = colorgful.MustApplyDefaultTheme(
		`${time:2006-01-02 15:04:05.000} ${level:%s:left:true} ${prefix}%s`,
		colorgful.Default,
	)
)

type (
	Level = lorg.Level
)

const (
	LevelFatal   = lorg.LevelFatal
	LevelError   = lorg.LevelError
	LevelWarning = lorg.LevelWarning
	LevelInfo    = lorg.LevelInfo
	LevelDebug   = lorg.LevelDebug
	LevelTrace   = lorg.LevelTrace
)

func init() {
	logger := lorg.NewLog()
	logger.SetIndentLines(true)

	if runtime.GOOS != "windows" {
		logger.SetFormat(theme)
		logger.SetOutput(theme)
	}

	if runtime.GOOS != "windows" {
		logger.SetShiftIndent(getShiftIndent(theme, ""))
	}

	implementation = cog.NewLogger(logger)

	proxy = Proxy{Logger: implementation}
	proxy.SetLevel(LevelInfo)
}

func SetLevel(level Level) {
	implementation.SetLevel(level)
}

func GetLevel() Level {
	return implementation.GetLevel()
}

func NewChild() Logger {
	return proxy.NewChild()
}

func GetImplementation() *cog.Logger {
	return proxy.Logger
}

func NewChildWithPrefix(prefix string) Logger {
	return proxy.NewChildWithPrefix(prefix)
}

func Fatalf(
	err error,
	message string,
	args ...interface{},
) {
	proxy.Fatalf(err, message, args...)
}

func Errorf(
	err error,
	message string,
	args ...interface{},
) {
	proxy.Errorf(err, message, args...)
}

func Warningf(
	err error,
	message string,
	args ...interface{},
) {
	proxy.Warningf(err, message, args...)
}

func Infof(
	context *karma.Context,
	message string,
	args ...interface{},
) {
	proxy.Infof(context, message, args...)
}

func Debugf(
	context *karma.Context,
	message string,
	args ...interface{},
) {
	proxy.Debugf(context, message, args...)
}

func Tracef(
	context *karma.Context,
	message string,
	args ...interface{},
) {
	proxy.Tracef(context, message, args...)
}

func Fatal(values ...interface{}) {
	proxy.Fatal(values...)
}

func Error(values ...interface{}) {
	proxy.Error(values...)
}

func Warning(values ...interface{}) {
	proxy.Warning(values...)
}

func Info(values ...interface{}) {
	proxy.Info(values...)
}

func Debug(values ...interface{}) {
	proxy.Debug(values...)
}

func Trace(values ...interface{}) {
	proxy.Trace(values...)
}

func getShiftIndent(theme *colorgful.Theme, prefix string) int {
	return len(
		regexp.MustCompile(`\x1b\[[^m]+m`).ReplaceAllString(
			fmt.Sprintf(theme.Render(lorg.LevelWarning, prefix), ""), "",
		),
	)
}
