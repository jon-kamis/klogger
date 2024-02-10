package klogger

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/internal/filelogger"
	"github.com/jon-kamis/klogger/pkg/loglevel"
	"github.com/stretchr/testify/assert"
)

const logLevelAllFileName = "properties\\test\\klogger-loglevel-all-properties.yml"
const logLevelErrorFileName = "properties\\test\\klogger-loglevel-error-properties.yml"

func TestEnter(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Enter(method)

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Info.String(), m[2])
	assert.Equal(t, method, m[3])
	assert.Equal(t, constants.Enter, strings.Trim(m[4], "\n"))

	//Test with this log level disabled
	os.Setenv("KloggerPropFileName", logLevelErrorFileName)
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	os.Setenv("KloggerPropFileName", logLevelErrorFileName)

	Enter(method)
	_, err = os.ReadFile("test-logs/application-test.log")
	assert.NotNil(t, err)

}

func TestExit(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Exit(method)

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Info.String(), m[2])
	assert.Equal(t, method, m[3])
	assert.Equal(t, constants.Exit, strings.Trim(m[4], "\n"))

	//Test with this log level disabled
	os.Setenv("KloggerPropFileName", logLevelErrorFileName)
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	Exit(method)
	_, err = os.ReadFile("test-logs/application-test.log")
	assert.NotNil(t, err)

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestInfo(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Info(method, "Testing info message with added messages: %s and %s", "m1", "m2")

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Info.String(), m[2])
	assert.Equal(t, method, m[3])

	//Test with this log level disabled
	os.Setenv("KloggerPropFileName", logLevelErrorFileName)
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	Info(method, "Testing info message with added messages: %s and %s", "m1", "m2")
	_, err = os.ReadFile("test-logs/application-test.log")
	assert.NotNil(t, err)

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestMultiLineLog(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Info(method, "Testing info message with added messages: %s\nand %s on a second line", "m1", "m2")

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	l := strings.Split(string(f), "\n")
	assert.Equal(t, 3, len(l))

	m := strings.Split(l[0], " ")
	assert.Equal(t, loglevel.Info.String(), m[2])
	assert.Equal(t, method, m[3])

	m1 := strings.Split(l[1], " ")
	assert.Equal(t, loglevel.Info.String(), m1[2])
	assert.Equal(t, method, m1[3])

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestDebug(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Debug(method, "Testing debug message with added messages: %s and %s", "m1", "m2")

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Debug.String(), m[2])
	assert.Equal(t, method, m[3])

	//Test with this log level disabled
	os.Setenv("KloggerPropFileName", logLevelErrorFileName)
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	Debug(method, "Testing info message with added messages: %s and %s", "m1", "m2")
	_, err = os.ReadFile("test-logs/application-test.log")
	assert.NotNil(t, err)

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestTrace(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Trace(method, "Testing trace message with added messages: %s and %s", "m1", "m2")

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Trace.String(), m[2])
	assert.Equal(t, method, m[3])

	//Test with this log level disabled
	os.Setenv("KloggerPropFileName", logLevelErrorFileName)
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	Trace(method, "Testing info message with added messages: %s and %s", "m1", "m2")
	_, err = os.ReadFile("test-logs/application-test.log")
	assert.NotNil(t, err)

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestWarn(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Warn(method, "Testing trace message with added messages: %s and %s", "m1", "m2")

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Warn.String(), m[2])
	assert.Equal(t, method, m[3])

	//Test with this log level disabled
	os.Setenv("KloggerPropFileName", logLevelErrorFileName)
	filelogger.CloseFile()
	os.RemoveAll("test-logs")
	Warn(method, "Testing info message with added messages: %s and %s", "m1", "m2")
	_, err = os.ReadFile("test-logs/application-test.log")
	assert.NotNil(t, err)

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestError(t *testing.T) {
	os.Setenv("KloggerPropFileName", logLevelAllFileName)
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestInfo"
	Error(method, "Testing error message with added messages: %s and %s", "m1", "m2")

	f, err := os.ReadFile("test-logs/application-test.log")
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Error.String(), m[2])
	assert.Equal(t, method, m[3])

	//Cleanup
	os.RemoveAll("test-logs")
}

func TestCheckFileRollover(t *testing.T) {
	os.Setenv("KloggerPropFileName", "properties\\test\\klogger-f-rollover-properties.yml")
	os.Setenv(constants.UseCacheEnvName, "false")
	filelogger.CloseFile()
	os.RemoveAll("test-logs")

	method := "TestCheckFileRollover"
	//Rollover is set to 10bytes. So running INFO twice will cause a file rollover to occur
	Info(method, "Testing info message with added messages: %s and %s", "m1", "m2")
	Info(method, "Testing info message with added messages: %s and %s", "m1", "m2")
	Info(method, "Testing info message with added messages: %s and %s", "m1", "m2")

	fn1 := fmt.Sprintf("test-logs/application-test_%v_1.log", time.Now().Format("2006-01-02"))
	fn2 := fmt.Sprintf("test-logs/application-test_%v_2.log", time.Now().Format("2006-01-02"))

	_, err := os.ReadFile(fn1)
	assert.Nil(t, err)

	f, err := os.ReadFile(fn2)
	assert.Nil(t, err)

	m := strings.Split(string(f), " ")

	assert.Equal(t, loglevel.Info.String(), m[2])
	assert.Equal(t, method, m[3])

	//Cleanup
	os.RemoveAll("test-logs")
}
