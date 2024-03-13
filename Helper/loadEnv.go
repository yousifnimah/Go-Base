package Helper

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
)

func LoadEnv() {
	err := godotenv.Load(filepath.Join(AppPath(), ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
