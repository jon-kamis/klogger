package config

import (
	"os"
	"sync/atomic"

	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/internal/properties"
	"github.com/jon-kamis/klogger/pkg/loglevel"
)

// Type KloggerConfig is a struct holding properties for the application
type KloggerConfig struct {
	PropFileName   string
	LogFileName    string
	LogFileDir     string
	DoRollover     bool
	DoSizeRollover bool
	RolloverSize   int64
	LogLevel       loglevel.LogLevel
	LogFileLevel   loglevel.LogLevel
	EnterLogLevel  loglevel.LogLevel
	ExitLogLevel   loglevel.LogLevel
}

var c atomic.Pointer[KloggerConfig] //Pointer cache

// Function GetConfig Caches and returns the config for the logger unless told not to do so for testing
func GetConfig() KloggerConfig {
	cached := c.Load()
	if cached != nil && os.Getenv(constants.UseCacheEnvName) != "false" {
		return *cached
	}

	config := loadConfig()
	cached = &config
	c.Store(cached)

	return *cached
}

// Function RefreshConfig causes Klogger module to wipe its cache and refresh its configuration
func RefreshConfig() KloggerConfig {

	config := loadConfig()
	cached := &config
	c.Store(cached)

	return *cached
}

func loadConfig() KloggerConfig {
	//Read in Config
	var config KloggerConfig

	//Load in properties
	props := properties.GetProperties()

	//Read in the rest of the props
	config.LogFileDir = properties.GetPropString(props.LogFileDir)
	config.LogFileName = properties.GetPropString(props.LogFileName)
	config.DoRollover = properties.GetPropBool(props.DoRollover)
	config.DoSizeRollover = properties.GetPropBool(props.DoSizeRollover)
	config.RolloverSize = int64(properties.GetPropInt(props.RolloverSize))
	config.LogFileLevel = properties.GetPropLogLevel(props.LogFileLevel)
	config.LogLevel = properties.GetPropLogLevel(props.LogLevel)
	config.EnterLogLevel = properties.GetPropLogLevel(props.EnterLogLevel)
	config.ExitLogLevel = properties.GetPropLogLevel(props.ExitLogLevel)

	return config
}
