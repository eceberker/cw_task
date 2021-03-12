package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

//GetEnv helps to get environmetn variables from .env file
func GetEnv() map[string]string {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Unable to read .env file")
	}
	return envs
}
