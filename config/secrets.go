package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SecretKey []byte

const KeySize = 32

func LoadSecrets() {
	_ = godotenv.Load(".env")
	SecretKey = []byte(os.Getenv("SWIFTTOKEN_SECRET"))
	if len(SecretKey) != KeySize {
		log.Fatal("SWIFT TOKEN SECRET environment variable is missing")
	}
}
