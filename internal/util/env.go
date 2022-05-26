package util

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {
	log.Println("Loading environment data")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment files")
	}
}
