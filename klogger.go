// Package logger contains custom logging methods
package klogger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jon-kamis/klogger/internal/config"
	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/pkg/loglevel"
)

// Function Enter returns a formated string used to declare where a method begins execution
// method - The method to write an enter log for
// l - The log levels to write to. If this is not set than the default log level for Enter logs is used
func Enter(method string, l ...loglevel.LogLevel) {

	if len(l) > 0 {
		for _, ll := range l {
			writeLog(constants.StdMsg, method, constants.Enter, ll)
		}
	} else {
		writeLog(constants.StdMsg, method, constants.Enter, loglevel.Info)
	}
}

// Function Exit returns a formated string used to declare where a method ends execution
// method - The method to write an enter log for
// l - The log levels to write to. If this is not set than the default log level for Enter logs is used
func Exit(method string, l ...loglevel.LogLevel) {
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
			writeLogToFile(l)
		}
	}
}

// Function WriteLogToFile writes a log to file based on config settings
// m - message to log
func writeLogToFile(msg string) {

	c := config.GetConfig()

	_ = os.Mkdir(c.LogFileDir, os.ModePerm)

	fn := fmt.Sprintf("%s/%s", c.LogFileDir, c.LogFileName)

	if c.DoRollover {
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
	fi, err := os.Stat(fmt.Sprintf("%s/%s", c.LogFileDir, c.LogFileName))

	if err == nil && c.DoSizeRollover && fi.Size() > c.RolloverSize {

		files, err := os.ReadDir(c.LogFileDir)

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
		ofn := fmt.Sprintf("%s/%s", c.LogFileDir, c.LogFileName)

		//New File Name
		fp := strings.Split(ofn, ".") //File Parts, 0 -> File Path and Name, 1 -> File Extension
		nfn := fmt.Sprintf("%s_%s_%d.%s", fp[0], dtStr, highestNum, fp[1])

		os.Rename(ofn, nfn)
	}
}
