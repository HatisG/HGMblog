package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("未找到.env文件")
	}
}

func Get(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
