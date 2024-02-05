package config

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	d := DefaultConfig
	//	o := "overridden value"
	fn := "\\properties\\test\\klogger-loadconf-properties.yml"
	fp, err := getFilePath()
	ffn := fmt.Sprintf("%s%s",fp,fn)
	assert.Nil(t, err)


	//First Assert that the default filepath is returned if the env variable is not set
	c := loadConfig()
	assert.Equal(t, d.PropFileName.Value, c.PropFileName.Value)

	//Second Assert that the env value overrides the default prop file name when set
	os.Setenv("PropFileName", ffn)
	c = loadConfig()
	assert.Equal(t, ffn, c.PropFileName.Value)

	//Test a value is loaded that was in the property file from the last test
	s, ok := c.LogFileName.Value.(string)
	assert.True(t, ok)
	assert.Equal(t, "app-conf-test.log", s)

	//Test a value that was not in the property file from the last test has its default value set
	s, ok = c.LogFileDir.Value.(string)
	assert.True(t, ok)
	assert.Equal(t, DefaultConfig.LogFileDir.Value, s)


}

func TestLoadFromEnvVariable(t *testing.T) {
	n := "TestName"
	v := "TestVal"
	vd := "TestValDefault"
	os.Setenv(n, v)

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

	fp, err := getFilePath()
	assert.Nil(t, err)

	//Test with property that exists in file
	r := loadProperty(p, fp+"\\properties\\test\\klogger-loadprop-properties.yml")
	s, ok := r.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, "test.log", s)

	//Test with a file that does not exist. Should return the same value as what is passed in
	r = loadProperty(p, fp+"\\properties\\test\\file-that-dne-properties.yml")
	s, ok = r.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, "test1", s)

	//Test with property that is not in file. Should return the same value as what is passed in
	p = Property{Name: "SomethingThatDoesNotExist", Value: "test2"}
	r = loadProperty(p, fp+"\\properties\\test\\klogger-loadprop-properties.yml")
	s, ok = r.Value.(string)

	assert.True(t, ok)
	assert.Equal(t, "test2", s)

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
