package globalconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) string {
	fmt.Println("Load .env file")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv(env)
	if apiKey == "" {
		fmt.Println("Set API_KEY environment variable.")
		return ""
	}

	return apiKey
}
