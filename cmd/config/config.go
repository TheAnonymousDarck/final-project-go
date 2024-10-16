package config

import (
	"os"
)

var (
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
)

func LoadConfig() {
	DBUser = getEnv("DB_USER", "root")
	DBPassword = getEnv("DB_PASSWORD", "")
	DBHost = getEnv("DB_HOST", "localhost")
	DBPort = getEnv("DB_PORT", "3306")
	DBName = getEnv("DB_NAME", "project-final-go")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
