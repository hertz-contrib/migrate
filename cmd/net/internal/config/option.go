package config

type Config struct {
	ServerVar string

	Addr         string
	IdleTimeout  string
	ReadTimeout  string
	WriteTimeout string
}

func NewConfig() *Config {
	return &Config{}
}
