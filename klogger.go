// Package logger contains custom logging methods
package klogger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jon-kamis/klogger/internal/config"
	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/internal/enum/loglevel"
)

var ll atomic.Pointer[loglevel.LogLevel]
var lfl atomic.Pointer[loglevel.LogLevel]

// Function Enter returns a formated string used to declare where a method begins execution
func Enter(method string) {
	writeLog(constants.StdMsg, method, constants.Enter, loglevel.Info)
}

// Function Exit returns a formated string used to declare where a method ends execution
func Exit(method string) {
	writeLog(constants.StdMsg, method, constants.Exit, loglevel.Info)
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

	//Check if anything will be logged by this command
	if logl < getLogLevel() && logl < getLogFileLevel() {
		return
	}

	//First fill in parameters
	lmsg := fmt.Sprintf(msg, args...)

	msgArr := strings.Split(lmsg, "\n")
	t := time.Now().Format(constants.TimeFormat)

	//Write to File if required
	if logl >= getLogLevel() {

		for _, m := range msgArr {
			l := fmt.Sprintf(mt, t, logl, me, m)
			fmt.Printf("%s\n", l)
		}

	}

	if logl >= getLogFileLevel() {
		for _, m := range msgArr {
			l := fmt.Sprintf(mt, t, logl, me, m)
			writeLogToFile(l)
		}
	}
}

// Function getLogLevel attempts to read in the log level from config
func getLogLevel() loglevel.LogLevel {

	cached := ll.Load()
	if cached != nil && os.Getenv(constants.UseCacheEnvName) != "false" {
		return *cached
	}

	logl, err := loglevel.GetLogLevelFromInterface(config.GetConfig().LogLevel.Value)

	if err != nil {
		panic("log level config is invalid!")
	}

	cached = &logl
	ll.Store(cached)

	return *cached
}

// Function getLogFileLevel attempts to read in the log file level from config
func getLogFileLevel() loglevel.LogLevel {

	cached := lfl.Load()
	if cached != nil && os.Getenv(constants.UseCacheEnvName) != "false" {
		return *cached
	}

	logfl, err := loglevel.GetLogLevelFromInterface(config.GetConfig().LogFileLevel.Value)

	if err != nil {
		panic("log level config is invalid!")
	}

	cached = &logfl
	lfl.Store(cached)

	return *cached
}

// Function WriteLogToFile writes a log to file based on config settings
// m - message to log
func writeLogToFile(msg string) {

	c := config.GetConfig()

	fd, ok := c.LogFileDir.Value.(string)

	if !ok {
		panic("log file directory is invalid!")
	}

	_ = os.Mkdir(fd, os.ModePerm)

	fn := fmt.Sprintf("%s/%s", fd, c.LogFileName.Value)

	if c.DoRollover.Value == true {
		checkFileRollover(c)
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("error occured %v\n", err)
		return
	}

	f.Write([]byte(msg + "\n"))
	f.Close()
}

// Function checkFileRollover determines if a file should be rolled over prior to writing to it
func checkFileRollover(c config.KloggerConfig) {
	fi, err := os.Stat(fmt.Sprintf("%s/%s", c.LogFileDir.Value, c.LogFileName.Value))

	rs, ok := c.RolloverSize.Value.(int)

	if err == nil && c.DoSizeRollover.Value == true && ok && fi.Size() > int64(rs) {
		d, ok := c.LogFileDir.Value.(string)

		if ok {
			files, err := os.ReadDir(d)

			if err != nil {
				return
			}

			dtStr := time.Now().Format("2006-01-02")
			highestNum := 0

			for _, file := range files {

				if strings.Contains(file.Name(), dtStr) {
					name := file.Name()
					name = name[:strings.IndexByte(name, '.')]
					num, err := strconv.Atoi(strings.Split(name, "_")[2])

					if err != nil {
						fmt.Printf("error occured: %v\n", err)
						return
					}

					if num > highestNum {
						highestNum = num
					}
				}
			}
			highestNum += 1

			//Original File Name
			ofn := fmt.Sprintf("%s/%s", c.LogFileDir.Value, c.LogFileName.Value)

			//New File Name
			fp := strings.Split(ofn, ".") //File Parts, 0 -> File Path and Name, 1 -> File Extension
			nfn := fmt.Sprintf("%s_%s_%d.%s", fp[0], dtStr, highestNum, fp[1])

			os.Rename(ofn, nfn)
		}
	}
}
