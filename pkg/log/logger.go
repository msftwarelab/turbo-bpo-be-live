package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	Prefix             = ""
	TimeFormat         = "2006-01-02 15:04:05"
	ShowDepth          bool
	DefaultCallerDepth = 2
	levelFlags         = []string{"verbose", "info", "warn", "error", "fatal"}
)

func Init(appname string) {
	Prefix = fmt.Sprintf("[%s]", appname)
}

type Level int

const (
	verbose Level = iota
	info
	warning
	errors
	fatal
)

func logger(level Level, depth int, format string, args ...interface{}) {
	var depthInfo string
	if ShowDepth {
		if depth == -1 {
			depth = DefaultCallerDepth
		}
		pc, file, line, ok := runtime.Caller(depth)
		if ok {
			fn := runtime.FuncForPC(pc)
			var fnName string
			if fn == nil {
				fnName = "?()"
			} else {
				fnName = strings.TrimLeft(filepath.Ext(fn.Name()), ".") + "()"
			}
			depthInfo = fmt.Sprintf("[%s:%d %s] ", filepath.Base(file), line, fnName)
		}
	}
	fmt.Printf("%s %s [%s] %s%s\n",
		Prefix, time.Now().Format(TimeFormat), levelFlags[level], depthInfo,
		fmt.Sprintf(format, args...))
	if level == fatal {
		os.Exit(1)
	}

}

func DebugD(depth int, format string, args ...interface{}) {
	logger(verbose, depth, format, args...)
}

func Debug(format string, args ...interface{}) {
	DebugD(-1, format, args...)
}

func WarnD(depth int, format string, args ...interface{}) {
	logger(warning, depth, format, args...)
}

func Warn(format string, args ...interface{}) {
	WarnD(-1, format, args...)
}

func InfoD(depth int, format string, args ...interface{}) {
	logger(info, depth, format, args...)
}

func Info(format string, args ...interface{}) {
	InfoD(-1, format, args...)
}

func ErrorD(depth int, format string, args ...interface{}) {
	logger(errors, depth, format, args...)
}

func Error(format string, args ...interface{}) {
	ErrorD(-1, format, args...)
}

func FatalD(depth int, format string, args ...interface{}) {
	logger(fatal, depth, format, args...)
}

func Fatal(format string, args ...interface{}) {
	FatalD(-1, format, args...)
}
