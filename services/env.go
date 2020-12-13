package services

import (
	"os"

	"github.com/joho/godotenv"
)

// GoDotEnvVariable Returns environment variable in .env file
func GoDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	return os.Getenv(key), nil
}
