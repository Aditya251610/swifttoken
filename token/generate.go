package token

import (
	"errors"

	"github.com/Aditya251610/swifttoken/crypto"
	"github.com/Aditya251610/swifttoken/encoder"
	"github.com/Aditya251610/swifttoken/types"
)

func GenerateToken(payload types.Payload) ([]byte, error) {
	if payload.IsValid() {
		encodedPayload, err := encoder.EncodePayload(payload)
		if err != nil {
			return nil, err
		}
		return crypto.Encrypt(encodedPayload)
	} else {
		return nil, errors.New("invalid payload")
	}
}
