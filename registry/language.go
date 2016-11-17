package registry

type Language struct {
	apply    func(LanguageOptions)
	versions func(string)
	preload  func(LanguageOptions)
}

type LanguageOptions struct {
	options map[string]string
}
