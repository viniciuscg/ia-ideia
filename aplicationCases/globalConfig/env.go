package globalconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)
a
func LoadEnv(env string) string {
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
