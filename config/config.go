package config

type Config struct {
	Router         RouterConfig
	DevicesService DevicesServiceConfig
}

type RouterConfig struct {
	Port string
}

type DevicesServiceConfig struct {
	UseMocks bool
}

func NewConfig() *Config {
	return &Config{
		Router: RouterConfig{
			Port: "8080",
		},

		DevicesService: DevicesServiceConfig{
			UseMocks: true,
		},
	}
}
