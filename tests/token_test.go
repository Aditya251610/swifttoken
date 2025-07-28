package tests

import (
	"testing"
	"time"

	"github.com/Aditya251610/swifttoken/config"
	"github.com/Aditya251610/swifttoken/token"
	"github.com/Aditya251610/swifttoken/types"
)

func init() {
	config.LoadSecrets()
}

func TestGenerateAndVerifyToken(t *testing.T) {
	// Step 1: Prepare a valid payload
	now := time.Now().Unix()
	payload := types.Payload{
		Sub:         "user123",
		IssuedAt:    now,
		ExpiresAt:   now + 1800, // 30 minutes later
		Nonce:       "test-nonce",
		SessionID:   "session-xyz",
		Permissions: []string{"read", "write"},
		Sliding:     false,
	}

	// Step 2: Generate the token
	tokenBytes, err := token.GenerateToken(payload)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Step 3: Verify the token
	decodedPayload, shouldRefresh, err := token.VerifyToken(tokenBytes)
	if err != nil {
		t.Fatalf("Failed to verify token: %v", err)
	}

	// Step 4: Assert all fields
	if decodedPayload.Sub != payload.Sub {
		t.Errorf("Expected Sub: %s, got: %s", payload.Sub, decodedPayload.Sub)
	}

	if decodedPayload.Nonce != payload.Nonce {
		t.Errorf("Expected Nonce: %s, got: %s", payload.Nonce, decodedPayload.Nonce)
	}

	if decodedPayload.SessionID != payload.SessionID {
		t.Errorf("Expected SessionID: %s, got: %s", payload.SessionID, decodedPayload.SessionID)
	}

	if decodedPayload.IssuedAt != payload.IssuedAt {
		t.Errorf("Expected IssuedAt: %d, got: %d", payload.IssuedAt, decodedPayload.IssuedAt)
	}

	if decodedPayload.ExpiresAt != payload.ExpiresAt {
		t.Errorf("Expected ExpiresAt: %d, got: %d", payload.ExpiresAt, decodedPayload.ExpiresAt)
	}

	if len(decodedPayload.Permissions) != len(payload.Permissions) {
		t.Errorf("Expected %d permissions, got: %d", len(payload.Permissions), len(decodedPayload.Permissions))
	}

	for i, perm := range payload.Permissions {
		if decodedPayload.Permissions[i] != perm {
			t.Errorf("Permission mismatch at index %d: expected %s, got %s", i, perm, decodedPayload.Permissions[i])
		}
	}

	if decodedPayload.Sliding != payload.Sliding {
		t.Errorf("Expected Sliding: %v, got: %v", payload.Sliding, decodedPayload.Sliding)
	}

	// Step 5: Ensure no refresh is needed
	if shouldRefresh {
		t.Errorf("Expected shouldRefresh to be false for non-sliding token")
	}
}
