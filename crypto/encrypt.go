package crypto

import (
	"crypto/rand"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"

	"github.com/Aditya251610/swifttoken/config"
)

const NonceSize = chacha20poly1305.NonceSizeX
const TagSize = 16

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

func Decrypt(token []byte) ([]byte, error) {
	if len(token) < NonceSize+TagSize {
		return nil, errors.New("token is too short")
	}

	nonce := token[0:24]
	ciphertext := token[24:]

	if config.SecretKey == nil || len(config.SecretKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid secret key")
	}

	cipher, err := chacha20poly1305.NewX(config.SecretKey)
	if err != nil {
		return nil, err
	}

	plaintext, e := cipher.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return nil, e
	}

	return plaintext, nil
}
