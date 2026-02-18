package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	Port         string
	FrontendURLs []string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	frontendURLsStr := getEnv("FRONTEND_URLS", "http://localhost:3000")
	frontendURLs := strings.Split(frontendURLsStr, ",")

	for i, url := range frontendURLs {
		frontendURLs[i] = strings.TrimSpace(url)
	}

	return &Config{
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "3306"),
		DBUser:       getEnv("DB_USER", "root"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "workshop_deck_2025"),
		JWTSecret:    getEnv("JWT_SECRET", "super-secret-jwt-lmao"),
		Port:         getEnv("PORT", "8080"),
		FrontendURLs: frontendURLs,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
