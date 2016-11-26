package configuration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"sync"

	"github.com/xLegoz/gum/registry"
	// "gopkg.in/yaml.v2"
  "smallfish/simpleyaml"
)

const filename = ".gumfile.yml"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Configuration struct {
	Language  registry.Handler
	Services  []registry.Handler
	Utilities []registry.Handler
	YAML      yamlConfiguration
}

type yamlConfiguration struct {
	Application map[string]interface{}
	Services    map[string]interface{}
	Utilities   map[string]interface{}
}

func (c *Configuration) LoadFromFile() error {
	path, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &c.YAML)
	return err
}

func (c *Configuration) LoadHandlers() error {
	var language string
	var languageVersion string
	for k, v := range c.YAML.Application["language"].(map[interface{}]interface{}) {
		language = k.(string)
		languageVersion = fmt.Sprintf("%f", v.(float64))
	}
	if _, ok := registry.Registry.Languages[language]; !ok {
		return fmt.Errorf("Language not registered: %s", language)
	}
	c.Language = registry.Registry.Languages[language]
	c.Language.Options = map[string]interface{}{"version": languageVersion}
	c.Services = []registry.Handler{}
	var handler registry.Handler
	for service, serviceOptions := range c.YAML.Services {
		if _, ok := registry.Registry.Services[service]; !ok {
			return fmt.Errorf("Service not registered: %s", service)
		}
		handler = registry.Registry.Services[service]
		handler.Options = serviceOptions.(map[string]interface{})
		c.Services = append(c.Services, handler)
	}
	c.Utilities = []registry.Handler{}
	for utility, utilityOptions := range c.YAML.Utilities {
		if _, ok := registry.Registry.Utilities[utility]; !ok {
			return fmt.Errorf("Utility not registered: %s", utility)
		}
		handler = registry.Registry.Utilities[utility]
		handler.Options = utilityOptions.(map[string]interface{})
		c.Utilities = append(c.Utilities, registry.Registry.Utilities[utility])
	}
	return nil
}

func (c *Configuration) ValidateConfiguration() error {
	// TODO(xLegoz): Validate configuration phase, i.e. check versions and structure for each handler
	return nil
}

type ByPriority []registry.Handler

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }
func (c *Configuration) PrioritizedHandlers() []registry.Handler {
	// TODO(xLegoz): cache this?
	handlers := append(c.Services, c.Language)
	handlers = append(handlers, c.Utilities...)
	sort.Sort(ByPriority(handlers))
	return handlers

}

type Action int

const (
	StopAction    Action = iota
	StartAction   Action = iota
	PrepareAction Action = iota
	CleanAction   Action = iota
)

func (c *Configuration) caller(action Action) error {
	var wg sync.WaitGroup
	var errors chan error
	for _, handler := range c.PrioritizedHandlers() {
		switch action {
		case StartAction:
			go handler.Start(handler.Options, &wg, errors)
		case PrepareAction:
			go handler.Prepare(handler.Options, &wg, errors)
		case StopAction:
			go handler.Stop(handler.Options, &wg, errors)
		case CleanAction:
			go handler.Clean(handler.Options, &wg, errors)
		default:
			panic("Invalid action for caller")

		}
	}
	select {
	case err := <-errors:
		return err
	default:
		wg.Done()
	}
	return nil
}

func (c *Configuration) Stop() error {
	err := c.caller(StopAction)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configuration) Clean() error {
	err := c.caller(CleanAction)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configuration) Start() error {
	err := c.caller(StartAction)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configuration) Prepare() error {
	err := c.caller(PrepareAction)
	if err != nil {
		return err
	}
	return nil
}

func (c *Configuration) LoadAndCheckConfiguration() error {
	err := c.LoadFromFile()
	if err != nil {
		return err
	}
	err = c.LoadHandlers()
	if err != nil {
		return err
	}
	err = c.ValidateConfiguration()
	fmt.Printf("Value: %#v\n", c)
	return err
}
