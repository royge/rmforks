package config

type Config interface {
	GetConfig(config string) error
}

type ConfigFunc func(config string) error

func (f ConfigFunc) GetConfig(config string) error {
	return f(config)
}
