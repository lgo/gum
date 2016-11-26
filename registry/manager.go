package registry

import "sync"

var Registry = CreateManager()

type Manager struct {
	Services  map[string]Handler
	Languages map[string]Handler
	Utilities map[string]Handler
}

func CreateManager() Manager {
	return Manager{
		Services:  make(map[string]Handler),
		Languages: make(map[string]Handler),
		Utilities: make(map[string]Handler),
	}
}

func RegisterService(handler Handler) {
	handler.Type = ServiceType
	Registry.Services[handler.Name] = handler
}

func RegisterLanguage(handler Handler) {
	handler.Type = LanguageType
	Registry.Languages[handler.Name] = handler
}

func RegisterUtility(handler Handler) {
	handler.Type = UtilityType
	Registry.Utilities[handler.Name] = handler
}

type HandlerType int

const (
	LanguageType HandlerType = iota
	ServiceType  HandlerType = iota
	UtilityType  HandlerType = iota
)

type handlerFunc func(map[string]interface{}, *sync.WaitGroup, chan error) error

func WrapHandler(fn func(map[string]interface{}) error) handlerFunc {
	return func(options map[string]interface{}, wg *sync.WaitGroup, errors chan error) error {
		defer wg.Done()
		var err error
		err = fn(options)
		if err != nil {
			errors <- err
			return err
		}
		return nil
	}
}

type Handler struct {
	Name     string
	Priority int
	Type     HandlerType
	Options  map[string]interface{}
	Start    handlerFunc
	Stop     handlerFunc
	Versions handlerFunc
	Prepare  handlerFunc
	Clean    handlerFunc
}
