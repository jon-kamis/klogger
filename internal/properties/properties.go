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
	PropFileName    Property
	LogFileName     Property
	LogFileDir      Property
	DoRollover      Property
	DoSizeRollover  Property
	DoDateRollover  Property
	RolloverSize    Property
	LogLevel        Property
	LogFileLevel    Property
	EnterLogLevel   Property
	ExitLogLevel    Property
	DoEnterExitLogs Property
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
	DoDateRollover: Property{
		Name:  constants.DoDateRollover,
		Value: constants.DefaultDoDateRolloverValue,
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
	DoEnterExitLogs: Property{
		Name:  constants.DoEnterExitLogs,
		Value: constants.DefaultDoEnterExitLogs,
	},
}

// Function GetProperties loads in all properties, first from env variables and then from a property file
func GetProperties() KloggerProperties {

	kp := DefaultProperties
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
	kp.LogFileDir = loadFromEnvVariable(kp.LogFileDir)
	kp.LogFileName = loadFromEnvVariable(kp.LogFileName)
	kp.DoRollover = loadFromEnvVariable(kp.DoRollover)
	kp.DoSizeRollover = loadFromEnvVariable(kp.DoSizeRollover)
	kp.DoDateRollover = loadFromEnvVariable(kp.DoDateRollover)
	kp.RolloverSize = loadFromEnvVariable(kp.RolloverSize)
	kp.LogFileLevel = loadFromEnvVariable(kp.LogFileLevel)
	kp.LogLevel = loadFromEnvVariable(kp.LogLevel)
	kp.EnterLogLevel = loadFromEnvVariable(kp.EnterLogLevel)
	kp.ExitLogLevel = loadFromEnvVariable(kp.ExitLogLevel)
	kp.DoEnterExitLogs = loadFromEnvVariable(kp.DoEnterExitLogs)

	//Next attempt to load each value from the property file if it exists
	if fExists {
		kp.LogFileDir = loadProperty(kp.LogFileDir, pfd)
		kp.LogFileName = loadProperty(kp.LogFileName, pfd)
		kp.DoRollover = loadProperty(kp.DoRollover, pfd)
		kp.DoDateRollover = loadProperty(kp.DoDateRollover, pfd)
		kp.DoSizeRollover = loadProperty(kp.DoSizeRollover, pfd)
		kp.RolloverSize = loadProperty(kp.RolloverSize, pfd)
		kp.LogFileLevel = loadProperty(kp.LogFileLevel, pfd)
		kp.LogLevel = loadProperty(kp.LogLevel, pfd)
		kp.EnterLogLevel = loadProperty(kp.EnterLogLevel, pfd)
		kp.ExitLogLevel = loadProperty(kp.ExitLogLevel, pfd)
		kp.DoEnterExitLogs = loadProperty(kp.DoEnterExitLogs, pfd)
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
		p.isLoaded = true
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
