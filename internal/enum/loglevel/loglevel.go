package loglevel

import "github.com/jon-kamis/klogger/internal/constants"

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
