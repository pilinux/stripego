package stripego

import "github.com/joho/godotenv"

// Env - load the configurations from .env
func Env() error {
	// load environment variables
	return godotenv.Load()
}
