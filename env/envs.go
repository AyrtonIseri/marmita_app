package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TwilioNumber string
}

var ENV = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		TwilioNumber: getEnv("TWILIO_NUMBER", "", true),
	}
}

func getEnv(key, fallback string, mandatory bool) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if mandatory {
		log.Fatalf("Server could not be initiallized. Missing mandatory ENV: %s", key)
	}

	return fallback
}
