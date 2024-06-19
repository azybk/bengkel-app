package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Get() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when load %s", err.Error())
	}

}