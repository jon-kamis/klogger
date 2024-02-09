package constants

import "github.com/jon-kamis/klogger/pkg/loglevel"

const PropFileName = "PropFileName"
const LogFileName = "LogFileName"
const LogFileDir = "LogFileDir"
const DoRollover = "DoRollover"
const DoSizeRollover = "DoSizeRollover"
const RolloverSize = "RolloverSize"
const LogLevel = "LogLevel"
const LogFileLevel = "LogFileLevel"
const EnterLogLevel = "EnterLogLevel"
const ExitLogLevel = "ExitLogLevel"

const EnvPrefix = "Klogger"

const DefaultPropFileValue = "properties/klogger-properties.yml"
const DefaultLogFileNameValue = "application.log"
const DefaultLogFileDirValue = "logs"
const DefaultDoRolloverValue = true
const DefaultDoSizeRolloverValue = true
const DefaultLogLevelValue = loglevel.Debug
const DefaultLogFileLevelValue = loglevel.Debug
const DefaultEnterLogLevelValue = loglevel.Info
const DefaultExitLogLevelValue = loglevel.Info

const TimeFormat = "2006-01-02 15:04:05"

// Default Byte Size to rollover file
const DefaultRolloverSize = 104857600

const Enter = "[ENTER]"
const Exit = "[EXIT]"
const StdMsg = "%v %s %s %s"

const UseCacheEnvName = "UseCache"
