package registry

var registry = CreateManager()

type Manager struct {
	services  map[string]Service
	languages map[string]Langauge
}

func CreateManager() {
	return Manager{}
}

func RegisterService(serviceName string, serviceHandler Service) {
	registry.services[serviceName] = serviceHandler
}

func RegisterLanguage(languageName string, languageHandler Service) {
	registry.languages[languageName] = languageHandler
}
