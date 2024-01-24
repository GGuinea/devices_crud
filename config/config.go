package config

type Config struct {
	Router RouterConfig
}


type RouterConfig struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Router: RouterConfig{
			Port: "8080",
		},
	}
}
