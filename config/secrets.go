package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var SecretKey []byte

const KeySize = 32

func LoadSecrets() error {
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../.env")

	envKey := os.Getenv("SWIFTTOKEN_SECRET")

	if len(envKey) != KeySize {
		fmt.Println("⚠️  Using fallback secret key")
		envKey = "12345678901234567890123456789012"
	} else {
		fmt.Println("✅ Loaded secret key from environment")
	}

	SecretKey = []byte(envKey)

	if len(SecretKey) != KeySize {
		return ErrInvalidSecretKey
	}

	return nil
}

var ErrInvalidSecretKey = &SecretKeyError{"SWIFTTOKEN_SECRET must be exactly 32 bytes"}

type SecretKeyError struct {
	Msg string
}

func (e *SecretKeyError) Error() string {
	return e.Msg
}
