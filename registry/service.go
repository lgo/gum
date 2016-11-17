package registry

type Service struct {
	apply    func(ServiceOptions)
	versions func(string)
}

type ServiceOptions struct {
	options map[string]string
}
