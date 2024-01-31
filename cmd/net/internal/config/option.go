package config

type HertzOption struct {
	Addr         string
	IdleTimeout  string
	ReadTimeout  string
	WriteTimeout string
}

func NewHertzOption() *HertzOption {
	return &HertzOption{}
}
