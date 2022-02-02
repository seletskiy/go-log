package log

import (
	"runtime"

	"github.com/reconquest/cog"
	"github.com/reconquest/karma-go"
)

type Proxy struct {
	*cog.Logger

	prefix string
}

func (proxy Proxy) SetLevel(level Level) {
	proxy.Logger.SetLevel(level)
}

func (proxy Proxy) GetLevel() Level {
	return proxy.GetLevel()
}

func (proxy Proxy) NewChild() Logger {
	return Proxy{
		Logger: proxy.Logger.NewChild(),
		prefix: proxy.prefix,
	}
}

func (proxy Proxy) NewChildWithPrefix(prefix string) Logger {
	if proxy.prefix != "" {
		prefix = proxy.prefix + " " + prefix
	}

	child := Proxy{
		Logger: proxy.Logger.NewChildWithPrefix(prefix),
		prefix: prefix,
	}

	if runtime.GOOS != "windows" {
		child.Logger.SetShiftIndent(getShiftIndent(theme, ""))
	}

	return child
}

func (proxy Proxy) Fatalf(err error, message string, args ...interface{}) {
	proxy.Logger.Fatalf(err, message, args...)
}

func (proxy Proxy) Errorf(err error, message string, args ...interface{}) {
	proxy.Logger.Errorf(err, message, args...)
}

func (proxy Proxy) Warningf(err error, message string, args ...interface{}) {
	proxy.Logger.Warningf(err, message, args...)
}

func (proxy Proxy) Infof(context *karma.Context, message string, args ...interface{}) {
	proxy.Logger.Infof(context, message, args...)
}

func (proxy Proxy) Debugf(context *karma.Context, message string, args ...interface{}) {
	proxy.Logger.Debugf(context, message, args...)
}

func (proxy Proxy) Tracef(context *karma.Context, message string, args ...interface{}) {
	proxy.Logger.Tracef(context, message, args...)
}

func (proxy Proxy) Fatal(values ...interface{})   { proxy.Logger.Fatal(values...) }
func (proxy Proxy) Error(values ...interface{})   { proxy.Logger.Error(values...) }
func (proxy Proxy) Warning(values ...interface{}) { proxy.Logger.Warning(values...) }
func (proxy Proxy) Info(values ...interface{})    { proxy.Logger.Info(values...) }
func (proxy Proxy) Debug(values ...interface{})   { proxy.Logger.Debug(values...) }
func (proxy Proxy) Trace(values ...interface{})   { proxy.Logger.Trace(values...) }
