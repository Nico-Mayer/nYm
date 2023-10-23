package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	PORT       string
	DB_URL     string
	PGDATABASE string
	PGHOST     string
	PGPASSWORD string
	PGPORT     string
	PGUSER     string
)

func init() {
	if os.Getenv("ENV") != "PROD" {
		fmt.Println("Loading local .env file")
		err := loadEnv()
		if err != nil {
			panic("Error loading .env file")
		}
	} else {
		fmt.Println("Loading Prod environment variables")
	}

	PORT = getEnv("PORT", "8080")
	DB_URL = getEnv("DATABASE_URL", "localhost:27017")
	PGDATABASE = getEnv("PGDATABASE", "postgres")
	PGHOST = getEnv("PGHOST", "localhost")
	PGPASSWORD = getEnv("PGPASSWORD", "postgres")
	PGPORT = getEnv("PGPORT", "5432")
	PGUSER = getEnv("PGUSER", "postgres")
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
		fmt.Println("Environment variable", key, "not set. Using fallback value:", fallback)
		return fallback
	}
	return value
}
