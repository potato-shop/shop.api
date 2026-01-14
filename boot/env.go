package boot

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file, using environment variables")
		return
	}

	log.Println("Loading .env file successfully")
}
