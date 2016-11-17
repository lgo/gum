package configuration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const filename = ".gumfile.yml"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Configuration struct {
	Application applicationConfiguration
}

type applicationConfiguration struct {
	Name         string
	Dependancies map[string]string
	Services     map[string]map[string]interface{}
	Proxies      map[string]string
}

func (c *Configuration) LoadConfiguration(path string) error {
	path, err := filepath.Abs(path + filename)
	if err != nil {
		return err
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return err
	}

	// TODO(xLegoz): Validate configuration phase, i.e. check versions and structure
	fmt.Printf("Value: %#v\n", c)
	return nil
}
