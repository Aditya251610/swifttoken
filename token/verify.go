package token

import (
	"errors"
	"time"

	"github.com/Aditya251610/swifttoken/crypto"
	"github.com/Aditya251610/swifttoken/encoder"
	"github.com/Aditya251610/swifttoken/types"
)

func VerifyToken(token []byte) (*types.Payload, bool, error) {
	currTime := time.Now().Unix()
	var shouldRefresh bool
	const slidingWindow = 15 * time.Minute
	decrypted, err := crypto.Decrypt(token)
	if err != nil {
		return nil, shouldRefresh, errors.New("failed to decrypt token")
	}

	payload, decodeErr := encoder.DecodePayload(decrypted)
	if decodeErr != nil {
		return nil, shouldRefresh, errors.New("failed to decode payload")
	}

	if !payload.IsValid() {
		return nil, shouldRefresh, errors.New("invalid payload structure")
	}

	if currTime < payload.IssuedAt || currTime > payload.ExpiresAt {
		return nil, shouldRefresh, errors.New("token is in future or expired")
	}

	if payload.Sliding && time.Until(time.Unix(payload.ExpiresAt, 0)) <= slidingWindow {
		shouldRefresh = true
	}

	return &payload, shouldRefresh, nil

}
