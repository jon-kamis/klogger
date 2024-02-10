package filelogger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jon-kamis/klogger/internal/config"
	"github.com/jon-kamis/klogger/internal/utils"
)

// var f is the file to write logs to. Note it will be closed automatically at program termination by the garbage collector
var f *os.File

func CloseFile() {
	if f != nil {
		f.Close()
		f = nil
	}
}

// Function WriteLogToFile writes a log to file based on config settings
// m - message to log
func WriteLogToFile(msg string) {

	c := config.GetConfig()

	_ = os.Mkdir(c.LogFileDir, os.ModePerm)

	if c.DoRollover {
		checkFileRollover(c)
	}

	if f == nil{

		fn := fmt.Sprintf("%s/%s", c.LogFileDir, c.LogFileName)

		var err error

		f, err = os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Printf("error occured %v\n", err)
			return
		}

	}

	f.Write([]byte(msg + "\n"))
	f.Sync()
}

// Function checkFileRollover determines if a file should be rolled over prior to writing to it
func checkFileRollover(c config.KloggerConfig) {
	
	var fi os.FileInfo
	var err error

	if f != nil {
		fi, err = f.Stat()
	} else {
		fi, err = os.Stat(fmt.Sprintf("%s/%s", c.LogFileDir, c.LogFileName))
	}

	if err != nil {
		return
	}

	//First check date Rollover
	if c.DoDateRollover && fi.ModTime().Before(utils.GetStartOfDay(time.Now())) {

		dtStr := fi.ModTime().Format("2006-01-02")
		renameFile(c, dtStr)

		//Return to prevent double rolling over
		return
	}

	if c.DoSizeRollover && fi.Size() > c.RolloverSize {
		renameFile(c, time.Now().Format("2006-01-02"))
	}
}

func renameFile(c config.KloggerConfig, s string) {
	//Original File Name
	ofn := fmt.Sprintf("%s/%s", c.LogFileDir, c.LogFileName)

	//New File Name
	fp := strings.Split(ofn, ".") //File Parts, 0 -> File Path and Name, 1 -> File Extension

	fnum := getHighestFileNumForDate(c, s)
	var nfn string

	if fnum >= 0 {
		nfn = fmt.Sprintf("%s_%s_%d.%s", fp[0], s, fnum+1, fp[1])
	} else {
		fmt.Printf("[Klogger] failed to rename file due to error")
		return
	}

	//Close the current file and set its value to nil. This will cause the next log to generate a new file
	f.Close()
	f = nil

	os.Rename(ofn, nfn)
}

// Function getHighestFileNumForDate returns the highest log file number for the given date
// c - the Klogger Config required for loading files
// s - a string representing the date to check for
func getHighestFileNumForDate(c config.KloggerConfig, s string) int {
	files, err := os.ReadDir(c.LogFileDir)

	if err != nil {
		panic("[Klogger] failed to read directory " + c.LogFileDir)
	}

	highestNum := 0

	for _, file := range files {

		if strings.Contains(file.Name(), s) {
			name := file.Name()
			name = name[:strings.IndexByte(name, '.')]
			num, err := strconv.Atoi(strings.Split(name, "_")[2])

			if err != nil {
				fmt.Printf("error occured: %v\n", err)
				return -1
			}

			if num > highestNum {
				highestNum = num
			}
		}
	}
	return highestNum
}
