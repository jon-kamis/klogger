package loglevel

import (
	"errors"

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
	None
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
	case None:
		return constants.LogLevelNone
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
	case 6:
		return None
	}
	return 0
}

// Function GetLogLevelFromInterface type asserts a log level from an interface and returns the value if it is valid. Otherwise causes a panic
func GetLogLevelFromInterface(i interface{}) (LogLevel, error) {
	ll, ok := i.(LogLevel)

	if !ok {
		return All, errors.New("log level is invalid")
	}

	return ll, nil
}
