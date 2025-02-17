package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	PORT         string
	DATABASE_URL string
)

func init() {
	if os.Getenv("ENV") != "PROD" {
		log.Println("Loading local .env file")
		err := loadEnv()
		if err != nil {
			panic("Error loading .env file")
		}
	} else {
		log.Println("Loading Prod environment variables")
	}

	PORT = getEnv("PORT", "8080")
	DATABASE_URL = getEnv("DATABASE_URL", "localhost:27017")
}

func loadEnv() error {
	envFile, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			err := os.Setenv(pair[0], pair[1])
			if err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Println("Environment variable", key, "not set. Using fallback value:", fallback)
		return fallback
	}
	return value
}
