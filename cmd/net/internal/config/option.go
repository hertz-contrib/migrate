package config

type Config struct {
	SrvVar string

	Addr         string
	IdleTimeout  string
	ReadTimeout  string
	WriteTimeout string
}

func NewConfig() *Config {
	return &Config{}
}
