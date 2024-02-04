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
	"github.com/jon-kamis/klogger/internal/enum/loglevel"
)

// Function WriteLogToFile writes a log to file based on config settings
// m - Method, l - Log Level, msg - Message
func writeLogToFile(l loglevel.LogLevel, msg string) {
	_ = os.Mkdir("logs", os.ModePerm)

	c := config.GetConfig()

	ll, ok := c.LogFileLevel.Value.(loglevel.LogLevel)

	if !ok {
		panic("log level config is invalid!")
	}

	if ll < l {
		return
	}

	fn := fmt.Sprintf("%s/%s", c.LogFileDir.Value, c.LogFileName.Value)

	if c.DoRollover.Value == true {
		checkFileRollover(c)
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("error occured %v\n", err)
		return
	}

	f.Write([]byte(msg))
	f.Close()
}

func checkFileRollover(c config.KloggerConfig) {
	fi, err := os.Stat(fmt.Sprintf("%s/%s", c.LogFileDir.Value, c.LogFileName.Value))

	rs, ok := c.RolloverSize.Value.(int64)

	if err == nil && c.DoSizeRollover.Value == true && ok && fi.Size() > rs {
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

// Function Enter returns a formated string used to declare where a method begins execution
func Enter(method string) {
	msg := fmt.Sprintf(constants.StdMsg, time.Now().Format(time.RFC3339), loglevel.Info, method, constants.Enter)
	fmt.Printf("%s\n", msg)
	writeLogToFile(loglevel.Info, constants.Enter)
}

// Function Exit returns a formated string used to declare where a method ends execution
func Exit(method string) {
	msg := fmt.Sprintf(constants.StdMsg, time.Now().Format(time.RFC3339), loglevel.Info, method, constants.Exit)
	fmt.Printf("%s\n", msg)
	writeLogToFile(loglevel.Info, constants.Exit)
}

// Function Error returns a formated string used to log a given error along with a custom error message and declaring which method the error occured in
func Error(method string, m string, args ...interface{}) {
	msg := fmt.Sprintf(fmt.Sprintf(constants.StdMsg, time.Now().Format(time.RFC3339), loglevel.Error, method, m), args...)
	fmt.Printf("%s\n", msg)
	writeLogToFile(loglevel.Error, msg)
}

// Function ExitError returns a formated string used to combine the Exit and Error functions together
func ExitError(method string, msg string, args ...interface{}) {
	Error(method, msg, args...)
	Exit(method)
}

// Fucntion Info returns a formatted string containing a custom message and the method that the message is coming from
func Info(method string, m string, args ...interface{}) {
	msg := fmt.Sprintf(fmt.Sprintf(constants.StdMsg, time.Now().Format(time.RFC3339), loglevel.Info, method, m), args...)
	fmt.Printf("%s\n", msg)
	writeLogToFile(loglevel.Info, constants.Exit)
}

// Fucntion Info returns a formatted string containing a custom message and the method that the message is coming from
func Debug(method string, m string, args ...interface{}) {
	msg := fmt.Sprintf(fmt.Sprintf(constants.StdMsg, time.Now().Format(time.RFC3339), loglevel.Debug, method, m), args...)
	fmt.Printf("%s\n", msg)
	writeLogToFile(loglevel.Debug, constants.Exit)
}

// Fucntion Info returns a formatted string containing a custom message and the method that the message is coming from
func Trace(method string, m string, args ...interface{}) {
	msg := fmt.Sprintf(fmt.Sprintf(constants.StdMsg, time.Now().Format(time.RFC3339), loglevel.Trace, method, m), args...)
	fmt.Printf("%s\n", msg)
	writeLogToFile(loglevel.Trace, constants.Exit)
}
