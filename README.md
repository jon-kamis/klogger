# klogger
## Overview

The klogger go module is designed to write simple yet informative logs to both stdout and a log file

By default, logs will be written to /logs/application.log

Log files will rollover by default each day or after reaching a reaching a size of 100Mb. The rollover file name will follow the format of `application_yyyy-mm-dd_i.log` where i is incremented starting at 1 for each date

## Methods

| Method | Description | Usage |
| :--- | :--- | :--- |
| Enter | Writes [Enter] to denote when a method is beginning. Can be written to multiple log levels if supplied, or if none are supplied writes to its default log level | klogger.Enter("method name", ...loglevel.LogLevel) |
| Exit | Writes [Exit] to denote when a method is ending. Can be written to multiple log levels if supplied, or if none are supplied writes to its default log level | klogger.Exit("method name", ...loglevel.LogLevel)
| Trace | Writes a log with Trace log level | Trace("method name", "message") |
| Debug | Writes a log with Debug log level | Debug("method name", "message") |
| Info | Writes a log with Info log level | Info("method name", "message") |
| Warn | Writes a log with Warning log level | Warn("method name", "message") |
| Error | Writes a log with Error log level | Error("method name", "message") |

## Properties
Multiple Properties exist that can be set with both a yaml property file and environment variables to modify how and when the module writes logs. A full list can be found below:

| Property | Type | Env Name | Default value | Description |
| :--- | :--- | :--- | :--- | :--- |
| PropFileName | string | KloggerPropFileName | properties/klogger-properties.yml | The property file to read values from |
| LogFileName | string | KloggerLogFileName | application.log | The name of the file to write logs to |
| LogFileDir | string | KloggerLogFileDir | logs | The directory to write log files in |
| DoRollover | bool | KloggerDoRollover | true | Determines whether to rollover log files |
| DoDateRollover | bool | KloggerDoDateRollover | true | Determines whether to rollover based on the current date |
| DoSizeRollover | bool | KloggerDoSizeRollover | true | Determines whether to rollover based on the size of the log file |
| RolloverSize | int | KloggerRolloverSize | 104857600 | The size limit in bytes for a log file to reach before rolling over |
| LogLevel | loglevel.LogLevel | KloggerLogLevel | 2 | The log level for stdout. Only logs above or equal to this value will be written. See [Log Levels](#log-levels) for more information |
| LogFileLevel | loglevel.LogLevel | KloggerLogFileLevel | 2 | The log level for log files. Only logs above or equal to this value will be written. See [Log Levels](#log-levels) for more information |
| EnterLogLevel | loglevel.LogLevel | KloggerEnterLogLevel | 2 | The log level to be used for ENTER logs. See [Log Levels](#log-levels) for more information |
| ExitLogLevel | loglevel.LogLevel | KloggerExitLogLevel | 2 | The log level to be used for EXIT logs. See [Log Levels](#log-levels) for more information |
| DoEnterExitLogs | bool | KloggerDoEnterExitLogs | true | Determines whether to write or ignore ENTER and EXIT logs |

Example Property file: 
```yaml
klogger:
  LogFileName: "application.log"
  LogFileDir: "logs"
  DoRollover: true
  DoDateRollover: true
  DoSizeRollover: true
  RolloverSize: 104857600
  LogLevel: 2
  LogFileLevel: 2
  EnterLogLevel: 3
  ExitLogLevel: 3
  DoEnterExitLogs: true

```

Another valid property file which only includes values to override

```yaml
klogger:
  LogFileName: "app.log"
  DoEnterExitLogs: false

```

## Log Levels

Log Levels are an enum type that can be set as integers in environment variables and property files. They can also be accessed externally in go.

| Log Level | GO Type | Integer Value |
| :--- | :--- | :--- |
| All | loglevel.All | 0 |
| Trace | loglevel.Trace | 1 |
| Debug | loglevel.Debug | 2 |
| Info | loglevel.Info | 3 |
| Warn | loglevel.Warn | 4 |
| Error | loglevel.Error | 5 |
| None | loglevel.None | 6 |
