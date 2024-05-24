package config

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBName     string
}

func initConfig() *Config {
	return &Config{
		PublicHost: "localhost",
		Port:       "3434",
	}
}
