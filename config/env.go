package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", ""),
		DBName:     getEnv("DB_NAME", "shop"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
