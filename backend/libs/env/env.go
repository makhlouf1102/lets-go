package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("all env variables loaded")
}

func Get(key string) string {
	return os.Getenv(key)
}
