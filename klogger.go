// Package logger contains custom logging methods
package klogger

import (
	"fmt"
	"strings"
	"time"

	"github.com/jon-kamis/klogger/internal/config"
	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/internal/filelogger"
	"github.com/jon-kamis/klogger/pkg/loglevel"
)

// Function Enter returns a formated string used to declare where a method begins execution
// method - The method to write an enter log for
// l - The log levels to write to. If this is not set than the default log level for Enter logs is used
// returns the time in which the log is written to track exit times if desired
func Enter(method string, l ...loglevel.LogLevel) time.Time {

	if !config.GetConfig().DoEnterExitLogs {
		return time.Now()
	}

	if len(l) > 0 {
		for _, ll := range l {
			writeLog(constants.StdMsg, method, constants.Enter, ll)
		}
	} else {
		writeLog(constants.StdMsg, method, constants.Enter, loglevel.Info)
	}

	return time.Now()
}

// Function Exit returns a formated string used to declare where a method ends execution
// method - The method to write an enter log for
// l - The log levels to write to. If this is not set than the default log level for Enter logs is used
func Exit(method string, l ...loglevel.LogLevel) {

	if !config.GetConfig().DoEnterExitLogs {
		return
	}

	if len(l) > 0 {
		for _, ll := range l {
			writeLog(constants.StdMsg, method, constants.Exit, ll)
		}
	} else {
		writeLog(constants.StdMsg, method, constants.Exit, loglevel.Info)
	}
}

// Function Error returns a formated string used to log a given error along with a custom error message and declaring which method the error occured in
func Error(method string, m string, args ...any) {
	writeLog(constants.StdMsg, method, m, loglevel.Error, args...)
}

// Function Warn returns a formated string used to log a given error along with a custom error message and declaring which method the warning occured in
func Warn(method string, m string, args ...any) {
	writeLog(constants.StdMsg, method, m, loglevel.Warn, args...)
}

// Function ExitError returns a formated string used to combine the Exit and Error functions together
func ExitError(method string, msg string, args ...any) {
	Error(method, msg, args...)
	Exit(method)
}

// Fucntion Info returns a formatted string containing a custom message and the method that the message is coming from
func Info(method string, m string, args ...any) {
	writeLog(constants.StdMsg, method, m, loglevel.Info, args...)
}

// Fucntion Debug returns a formatted string containing a custom message and the method that the message is coming from
func Debug(method string, m string, args ...any) {
	writeLog(constants.StdMsg, method, m, loglevel.Debug, args...)
}

// Function Trace returns a formatted string containing a custom message and the method that the message is coming from
func Trace(method string, m string, args ...any) {
	writeLog(constants.StdMsg, method, m, loglevel.Trace, args...)
}

// Function RefreshConfig causes the Klogger module to refresh its config
func RefreshConfig() {
	config.RefreshConfig()
}

// Function writeLog writes a log to stdout and a log file
// mt - message template
// m - method
// msg - message to log
func writeLog(mt string, me string, msg string, logl loglevel.LogLevel, args ...any) {

	c := config.GetConfig()

	//Check if anything will be logged by this command
	if logl < c.LogLevel && logl < c.LogFileLevel {
		return
	}

	//First fill in parameters
	lmsg := fmt.Sprintf(msg, args...)

	msgArr := strings.Split(lmsg, "\n")
	t := time.Now().Format(constants.TimeFormat)

	//Write to File if required
	if logl >= config.GetConfig().LogLevel {

		for _, m := range msgArr {
			l := fmt.Sprintf(mt, t, logl, me, m)
			fmt.Printf("%s\n", l)
		}

	}

	if logl >= config.GetConfig().LogFileLevel {
		for _, m := range msgArr {
			l := fmt.Sprintf(mt, t, logl, me, m)
			filelogger.WriteLogToFile(l)
		}
	}
}
