package config

import (
	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/internal/enum/loglevel"
)

// Type Property is a pairing of property names and their default values
type Property struct {
	Name  string
	Value interface{}
}

// Type PropertyFileData is a struct representing a property file data for this go module
type Config struct {
	Klogger map[string]interface{}
}

// Type KloggerConfig is a struct holding properties for the application
type KloggerConfig struct {
	PropFileName   Property
	LogFileName    Property
	LogFileDir     Property
	DoRollover     Property
	DoSizeRollover Property
	RolloverSize   Property
	LogLevel       Property
	LogFileLevel   Property
}

// Function GetDefaultConfig returns a FinanceManagerConfig object containing the default values for each environment variable
var DefaultConfig = KloggerConfig{
	PropFileName: Property{
		Name:  "PropFileName",
		Value: constants.DefaultPropFileName,
	},
	LogFileName: Property{
		Name:  "LogFileName",
		Value: constants.DefaultLogFileName,
	},
	LogFileDir: Property{
		Name:  "LogFileDir",
		Value: constants.DefaultLogFileDir,
	},
	DoRollover: Property{
		Name:  "DoRollover",
		Value: constants.DefaultDoRollover,
	},
	DoSizeRollover: Property{
		Name:  "DoSizeRollover",
		Value: constants.DefaultDoSizeRollover,
	},
	RolloverSize: Property{
		Name:  "RolloverSize",
		Value: constants.DefaultRolloverSize,
	},
	LogLevel: Property{
		Name:  "LogLevel",
		Value: loglevel.Debug,
	},
	LogFileLevel: Property{
		Name:  "LogFileLevel",
		Value: loglevel.Debug,
	},
}
