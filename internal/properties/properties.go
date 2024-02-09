package properties

import (
	"fmt"
	"os"

	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/pkg/loglevel"
	"gopkg.in/yaml.v2"
)

// Type PropertyFileData is a struct representing a property file data for this go module
type PropertyFileData struct {
	Klogger map[string]interface{}
}

// Type Property is a pairing of property names and their default values
type Property struct {
	Name     string
	Value    interface{}
	isLoaded bool
}

// Type KloggerProperties is a struct holding properties for the application
type KloggerProperties struct {
	PropFileName   Property
	LogFileName    Property
	LogFileDir     Property
	DoRollover     Property
	DoSizeRollover Property
	RolloverSize   Property
	LogLevel       Property
	LogFileLevel   Property
	EnterLogLevel  Property
	ExitLogLevel   Property
}

type Number interface {
	int | int64
}

// Variable DefaultProperties holds the default application properties
var DefaultProperties = KloggerProperties{
	PropFileName: Property{
		Name:  constants.PropFileName,
		Value: constants.DefaultPropFileValue,
	},
	LogFileName: Property{
		Name:  constants.LogFileName,
		Value: constants.DefaultLogFileNameValue,
	},
	LogFileDir: Property{
		Name:  constants.LogFileDir,
		Value: constants.DefaultLogFileDirValue,
	},
	DoRollover: Property{
		Name:  constants.DoRollover,
		Value: constants.DefaultDoRolloverValue,
	},
	DoSizeRollover: Property{
		Name:  constants.DoSizeRollover,
		Value: constants.DefaultDoSizeRolloverValue,
	},
	RolloverSize: Property{
		Name:  constants.RolloverSize,
		Value: constants.DefaultRolloverSize,
	},
	LogLevel: Property{
		Name:  constants.LogLevel,
		Value: constants.DefaultLogLevelValue,
	},
	LogFileLevel: Property{
		Name:  constants.LogFileLevel,
		Value: constants.DefaultLogFileLevelValue,
	},
	EnterLogLevel: Property{
		Name:  constants.EnterLogLevel,
		Value: constants.DefaultEnterLogLevelValue,
	},
	ExitLogLevel: Property{
		Name:  constants.ExitLogLevel,
		Value: constants.DefaultExitLogLevelValue,
	},
}

// Function GetProperties loads in all properties, first from env variables and then from a property file
func GetProperties() KloggerProperties {

	var kp KloggerProperties
	var pfd PropertyFileData
	fExists := true

	//First Check env variable for property file name
	kp.PropFileName = loadFromEnvVariable(DefaultProperties.PropFileName)

	fn, ok := kp.PropFileName.Value.(string)

	if ok {
		//Attempt to read the property file
		pf, err := os.ReadFile(fn)

		if err != nil {
			fmt.Printf("error reading property file: %v\n", err)
			fExists = false

		} else {

			err = yaml.Unmarshal(pf, &pfd)

			if err != nil {
				fmt.Printf("failed to unmarshal yaml: %v\n", err)
				fExists = false
			}
		}

	} else {
		panic("property file name is invalid")
	}

	//First attempt to load each value from the environment
	kp.LogFileDir = loadFromEnvVariable(DefaultProperties.LogFileDir)
	kp.LogFileName = loadFromEnvVariable(DefaultProperties.LogFileName)
	kp.DoRollover = loadFromEnvVariable(DefaultProperties.DoRollover)
	kp.DoSizeRollover = loadFromEnvVariable(DefaultProperties.DoSizeRollover)
	kp.RolloverSize = loadFromEnvVariable(DefaultProperties.RolloverSize)
	kp.LogFileLevel = loadFromEnvVariable(DefaultProperties.LogFileLevel)
	kp.LogLevel = loadFromEnvVariable(DefaultProperties.LogLevel)
	kp.EnterLogLevel = loadFromEnvVariable(DefaultProperties.EnterLogLevel)
	kp.ExitLogLevel = loadFromEnvVariable(DefaultProperties.ExitLogLevel)

	//Next attempt to load each value from the property file if it exists
	if fExists {
		kp.LogFileDir = loadProperty(DefaultProperties.LogFileDir, pfd)
		kp.LogFileName = loadProperty(DefaultProperties.LogFileName, pfd)
		kp.DoRollover = loadProperty(DefaultProperties.DoRollover, pfd)
		kp.DoSizeRollover = loadProperty(DefaultProperties.DoSizeRollover, pfd)
		kp.RolloverSize = loadProperty(DefaultProperties.RolloverSize, pfd)
		kp.LogFileLevel = loadProperty(DefaultProperties.LogFileLevel, pfd)
		kp.LogLevel = loadProperty(DefaultProperties.LogLevel, pfd)
		kp.EnterLogLevel = loadProperty(DefaultProperties.EnterLogLevel, pfd)
		kp.ExitLogLevel = loadProperty(DefaultProperties.ExitLogLevel, pfd)
	}

	return kp
}

// Function loadFromEnvVariable attempts to update a single property value from an environment variable
func loadFromEnvVariable(p Property) Property {

	//Attempt to get the environment variable
	v := os.Getenv(constants.EnvPrefix + p.Name)

	if v != "" {
		fmt.Printf("[Klogger] Read environment variable for property: %s\n", p.Name)
		p.Value = v
		p.isLoaded = true
	}

	return p
}

// Function loadProperty attempts to Load a single property from a property object if it has not already been loaded
func loadProperty(p Property, pfd PropertyFileData) Property {

	if !p.isLoaded && pfd.Klogger[p.Name] != nil {

		fmt.Printf("[Klogger] Read property file entry for property: %s\n", p.Name)
		p.Value = pfd.Klogger[p.Name]
	}

	return p
}

func GetPropString(p Property) string {
	s, ok := p.Value.(string)
	validateValue(p.Name, ok)

	return s
}

func GetPropBool(p Property) bool {
	b, ok := p.Value.(bool)
	validateValue(p.Name, ok)

	return b
}

func GetPropInt(p Property) int {
	i, ok := p.Value.(int)
	validateValue(p.Name, ok)

	return i
}

func GetPropLogLevel(p Property) loglevel.LogLevel {

	//First check if allready typed to LogLevel
	ll, ok := p.Value.(loglevel.LogLevel)
	
	if ok {
		return ll
	}

	lli, ok := p.Value.(int)

	validateValue(p.Name, ok)

	return loglevel.GetLogLevel(lli)
}

func validateValue(n string, isValid bool) {
	if !isValid {
		panic(fmt.Sprintf("property for %s is invalid", n))
	}
}
