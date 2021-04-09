package tests

// AppOptionsProvider provides options for the app.
type AppOptionsProvider struct {
	storage map[string]interface{}
}

// Set sets app option.
func (ao AppOptionsProvider) Set(key string, value interface{}) {
	ao.storage[key] = value
}

// Get implements AppOptions.
func (ao AppOptionsProvider) Get(key string) interface{} {
	return ao.storage[key]
}

// NewAppOptionsProvider creates a new AppOptionsProvider.
func NewAppOptionsProvider() AppOptionsProvider {
	return AppOptionsProvider{
		storage: make(map[string]interface{}),
	}
}
