package config

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/internal/enum/loglevel"
	"gopkg.in/yaml.v2"
)

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

func loadConfig() KloggerConfig {
	//Read in Config
	var config KloggerConfig

	//First Check env variable for property file name
	config.PropFileName = loadFromEnvVariable(DefaultConfig.PropFileName)

	fn, ok := config.PropFileName.Value.(string)

	if ok {

		//Read in the rest of the props
		config.LogFileDir = loadProperty(DefaultConfig.LogFileDir, fn)
		config.LogFileName = loadProperty(DefaultConfig.LogFileName, fn)
		config.DoRollover = loadProperty(DefaultConfig.DoRollover, fn)
		config.DoSizeRollover = loadProperty(DefaultConfig.DoSizeRollover, fn)
		config.RolloverSize = loadProperty(DefaultConfig.RolloverSize, fn)
		config.LogFileLevel = loadProperty(DefaultConfig.LogFileLevel, fn)
		config.LogLevel = loadProperty(DefaultConfig.LogLevel, fn)

	} else {
		panic("property file name is invalid")
	}

	return config
}

// Function loadFromEnvVariable attempts to load a single property from an environment variable
// If the environment variable is not set it returns the value from p
func loadFromEnvVariable(p Property) Property {
	//First try to load from environment variable
	value := os.Getenv(p.Name)

	r := Property{
		Name:  p.Name,
		Value: p.Value,
	}

	if value != "" {
		fmt.Printf("[Klogger] Read environment variable for property: %s", p.Name)
		r.Value = value
	}

	return r
}

// Function loadProperty attempts to Load a single property from a property file.
// If the file does not exist or does not contain the given property, then the default value is used
func loadProperty(p Property, fn string) Property {

	propFile, err := os.ReadFile(fn)

	r := p

	if err != nil {
		fmt.Printf("error reading property file: %v\n", err)
	} else {
		var props Config
		err = yaml.Unmarshal(propFile, &props)

		if err != nil {
			fmt.Printf("failed to unmarshal yaml: %v\n", err)
		} else if props.Klogger[p.Name] != nil {

			if p.Name == DefaultConfig.LogLevel.Name || p.Name == DefaultConfig.LogFileLevel.Name {
				i, ok := props.Klogger[p.Name].(int)

				if !ok {
					fmt.Printf("Invalid value of %v for property: %s\n", props.Klogger[p.Name], p.Name)
					return p
				}

				fmt.Printf("[Klogger] Read property file entry for property: %s", p.Name)
				r.Value = loglevel.GetLogLevel(int64(i))
				return r
			}
			fmt.Printf("[Klogger] Read property file entry for property: %s", p.Name)
			r.Value = props.Klogger[p.Name]
			return r

		}
	}

	return p
}
