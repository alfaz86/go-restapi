package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host string
	Port string
	URL  string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

var Cfg AppConfig

func LoadConfig() {
	_ = godotenv.Load()

	Cfg = AppConfig{
		Host: getEnv("APP_HOST", "0.0.0.0"),
		Port: getEnv("APP_PORT", "8080"),
		URL:  getEnv("APP_URL", "http://localhost:8080/"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASSWORD", "postgres"),
		DBName: getEnv("DB_NAME", "go_restapi"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func (c AppConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c AppConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort,
	)
}
