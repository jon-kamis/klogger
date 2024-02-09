package properties

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jon-kamis/klogger/internal/constants"
	"github.com/jon-kamis/klogger/pkg/loglevel"
	"github.com/stretchr/testify/assert"
)

func TestGetProperties(t *testing.T) {
	d := DefaultProperties
	//	o := "overridden value"

	fn := "\\properties\\test\\klogger-loadconf-properties.yml"
	fp, err := getFilePath()
	fp = strings.Replace(fp, "\\internal\\properties", "", -1)
	ffn := fmt.Sprintf("%s%s", fp, fn)
	assert.Nil(t, err)

	//First Assert that the default filepath is returned if the env variable is not set
	c := GetProperties()
	assert.Equal(t, d.PropFileName.Value, c.PropFileName.Value)

	//Second Assert that the env value overrides the default prop file name when set
	os.Setenv("KloggerPropFileName", ffn)
	c = GetProperties()
	assert.Equal(t, ffn, c.PropFileName.Value)

	//Test a value is loaded that was in the property file from the last test
	s, ok := c.LogFileName.Value.(string)
	assert.True(t, ok)
	assert.Equal(t, "app-conf-test.log", s)

	//Test a value that was not in the property file from the last test has its default value set
	s, ok = c.LogFileDir.Value.(string)
	assert.True(t, ok)
	assert.Equal(t, DefaultProperties.LogFileDir.Value, s)

}

func TestLoadFromEnvVariable(t *testing.T) {
	n := "TestName"
	v := "TestVal"
	vd := "TestValDefault"
	os.Setenv(constants.EnvPrefix+n, v)

	p := Property{
		Name:  n,
		Value: vd,
	}

	p1 := loadFromEnvVariable(p)

	s, ok := p1.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, v, s)
}

func TestLoadProperty(t *testing.T) {

	p := Property{
		Name:  "LogFileName",
		Value: "test1",
	}

	d := make(map[string]interface{})
	d[p.Name] = "test.log"

	pfd := PropertyFileData{
		Klogger: d,
	}

	//Test with property that exists in property file data
	r := loadProperty(p, pfd)
	s, ok := r.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, "test.log", s)

	//Test with property that exists in property file data but has bene loaded. Should return the same valueu as what is passed in
	p.isLoaded = true
	r = loadProperty(p, pfd)
	s, ok = r.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, p.Value, s)

	//Test with property that is not in file. Should return the same value as what is passed in
	p = Property{Name: "SomethingThatDoesNotExist", Value: "test2"}
	r = loadProperty(p, pfd)
	s, ok = r.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, "test2", s)

}

func TestGetPropString(t *testing.T) {
	v := "string"
	p := Property{
		Name:  "prop",
		Value: v,
	}

	s := GetPropString(p)
	assert.Equal(t, v, s)

	p.Value = 1

	assert.Panics(t, func() { GetPropString(p) })

}

func TestGetPropBool(t *testing.T) {
	v := true
	p := Property{
		Name:  "prop",
		Value: v,
	}

	s := GetPropBool(p)
	assert.Equal(t, v, s)

	p.Value = 1

	assert.Panics(t, func() { GetPropBool(p) })

}

func TestGetPropInt(t *testing.T) {
	v := 1
	p := Property{
		Name:  "prop",
		Value: v,
	}

	s := GetPropInt(p)
	assert.Equal(t, v, s)

	p.Value = "str"

	assert.Panics(t, func() { GetPropInt(p) })

}

func TestGetPropLogLevel(t *testing.T) {
	v := loglevel.Debug
	p := Property{
		Name:  "prop",
		Value: v,
	}

	s := GetPropLogLevel(p)
	assert.Equal(t, v, s)

	//Also accepts integers

	v = 1
	p.Value = v
	s = GetPropLogLevel(p)
	assert.Equal(t, v, s)

	p.Value = "str"
	assert.Panics(t, func() { GetPropLogLevel(p) })

}

func getFilePath() (string, error) {
	fp, err := os.Getwd()

	if err != nil {
		return "", err
	}

	//Do this to handle both windows and linux file systems
	fp = strings.Replace(fp, "internal\\config", "", -1)
	fp = strings.Replace(fp, "internal/config", "", -1)

	return fp, nil
}
