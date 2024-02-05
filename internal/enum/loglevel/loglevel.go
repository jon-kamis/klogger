package loglevel

import (
	"github.com/jon-kamis/klogger/internal/constants"
)

type LogLevel int64

const (
	All LogLevel = iota
	Trace
	Debug
	Info
	Warn
	Error
)

func (l LogLevel) String() string {
	switch l {
	case Trace:
		return constants.LogLevelTrace
	case Debug:
		return constants.LogLevelDebug
	case Info:
		return constants.LogLevelInfo
	case Warn:
		return constants.LogLevelWarn
	case Error:
		return constants.LogLevelErr
	}
	return "UNKWN"
}

func GetLogLevel(i int64) LogLevel {
	switch i {
	case 0:
		return All
	case 1:
		return Trace
	case 2:
		return Debug
	case 3:
		return Info
	case 4:
		return Warn
	case 5:
		return Error
	}
	return 0
}
