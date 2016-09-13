package config

type Config interface {
	GetConfig(config string) (Config, error)
}

type Auth interface {
	GetCredentials(config Config) (interface{}, error)
}

type ConfigFunc func(config string) (Config, error)
type AuthFunc func(config Config) (interface{}, error)

func (f ConfigFunc) GetConfig(config string) (Config, error) {
	return f(config)
}

func (f AuthFunc) GetCredentials(config Config) (interface{}, error) {
	return f(config)
}
