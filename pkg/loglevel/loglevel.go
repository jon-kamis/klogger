package loglevel

import (
	"errors"
)

// Type LogLevel is an int enum used to determine which logs should and should not be written by a configuration, as well as what value to log for each level
type LogLevel int

const (
	All LogLevel = iota
	Trace
	Debug
	Info
	Warn
	Error
	None
)

const logLevelTrace = "TRACE"
const logLevelDebug = "DEBUG"
const logLevelInfo = "INFO"
const logLevelWarn = "WARN"
const logLevelErr = "ERROR"
const logLevelNone = "NONE"

// Function String is used when printing a LogLevel Object
func (l LogLevel) String() string {
	switch l {
	case Trace:
		return logLevelTrace
	case Debug:
		return logLevelDebug
	case Info:
		return logLevelInfo
	case Warn:
		return logLevelWarn
	case Error:
		return logLevelErr
	case None:
		return logLevelNone
	}
	return "UNKWN"
}

// Function GetLogLevel accepts an int argument and returns the corresponding LogLevel for that value if one exists or defaults to All
func GetLogLevel(i int) LogLevel {
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
