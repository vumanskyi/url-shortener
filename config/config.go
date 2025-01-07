package config

type AppConfig struct {
	Host string
	Port string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Host: "localhost",
		Port: "3000",
	}
}
