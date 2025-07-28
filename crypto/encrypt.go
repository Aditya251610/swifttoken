package crypto

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/Aditya251610/swifttoken/config"
)

const NonceSize = chacha20poly1305.NonceSizeX

func Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, NonceSize)
	if config.SecretKey == nil || len(config.SecretKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid secret key")
	}
	cipher, err := chacha20poly1305.NewX(config.SecretKey)
	if err != nil {
		return nil, err
	}
	i, nonceErr := rand.Read(nonce)
	if nonceErr != nil || i != NonceSize {
		return nil, errors.New("failed to generate nonce")
	}

	ciphertext := cipher.Seal(nil, nonce, plaintext, nil)
	finalOutput := append(nonce, ciphertext...)
	return finalOutput, nil
}
