package tests

import (
	"testing"
	"time"

	"github.com/Aditya251610/swifttoken/config"
	"github.com/Aditya251610/swifttoken/token"
	"github.com/Aditya251610/swifttoken/types"
)

func init() {
	_ = config.LoadSecrets() // Will use fallback if .env is not set
}

// Helper to generate valid payloads
func validPayload(sliding bool, expiresOffset int64) types.Payload {
	now := time.Now().Unix()
	return types.Payload{
		Sub:         "user123",
		IssuedAt:    now,
		ExpiresAt:   now + expiresOffset,
		Nonce:       "test-nonce",
		SessionID:   "session-xyz",
		Permissions: []string{"read", "write"},
		Sliding:     sliding,
	}
}

// ✅ Basic working test
func TestGenerateAndVerifyToken(t *testing.T) {
	payload := validPayload(false, 1800)
	tokenBytes, err := token.GenerateToken(payload)
	if err != nil {
		t.Fatalf("GenerateToken failed: %v", err)
	}

	decodedPayload, shouldRefresh, err := token.VerifyToken(tokenBytes)
	if err != nil {
		t.Fatalf("VerifyToken failed: %v", err)
	}

	if decodedPayload.Sub != payload.Sub {
		t.Errorf("Sub mismatch: expected %s, got %s", payload.Sub, decodedPayload.Sub)
	}

	if shouldRefresh {
		t.Errorf("Expected shouldRefresh to be false")
	}
}

// 🔁 Sliding token nearing expiry should trigger refresh
func TestSlidingTokenTriggersRefresh(t *testing.T) {
	payload := validPayload(true, 600)
	tokenBytes, _ := token.GenerateToken(payload)
	_, shouldRefresh, _ := token.VerifyToken(tokenBytes)
	if !shouldRefresh {
		t.Errorf("Expected shouldRefresh to be true for sliding token")
	}
}

// 🔁 Sliding token far from expiry should not trigger refresh
func TestSlidingTokenNoRefresh(t *testing.T) {
	payload := validPayload(true, 3600)
	tokenBytes, _ := token.GenerateToken(payload)
	_, shouldRefresh, _ := token.VerifyToken(tokenBytes)
	if shouldRefresh {
		t.Errorf("Expected shouldRefresh to be false")
	}
}

// ❌ Expired token
func TestExpiredToken(t *testing.T) {
	payload := validPayload(false, -60)
	tokenBytes, _ := token.GenerateToken(payload)
	_, _, err := token.VerifyToken(tokenBytes)
	if err == nil {
		t.Errorf("Expected error for expired token")
	}
}

// ❌ Future-issued token
func TestFutureToken(t *testing.T) {
	now := time.Now().Unix()
	payload := types.Payload{
		Sub:         "future-user",
		IssuedAt:    now + 1000,
		ExpiresAt:   now + 2000,
		Nonce:       "nonce",
		SessionID:   "sid",
		Permissions: []string{"read"},
		Sliding:     false,
	}
	tokenBytes, _ := token.GenerateToken(payload)
	_, _, err := token.VerifyToken(tokenBytes)
	if err == nil {
		t.Errorf("Expected error for future-issued token")
	}
}

// ❌ Empty payload
func TestEmptyFieldsPayload(t *testing.T) {
	payload := types.Payload{}
	_, err := token.GenerateToken(payload)
	if err == nil {
		t.Errorf("Expected error for empty payload")
	}
}

// ❌ Invalid timestamps
func TestInvalidTimestamps(t *testing.T) {
	now := time.Now().Unix()
	payload := types.Payload{
		Sub:         "abc",
		IssuedAt:    now,
		ExpiresAt:   now - 100,
		Nonce:       "nonce",
		SessionID:   "sid",
		Permissions: []string{"admin"},
	}
	_, err := token.GenerateToken(payload)
	if err == nil {
		t.Errorf("Expected error for ExpiresAt < IssuedAt")
	}
}

// 🧪 Tampered token
func TestTamperedToken(t *testing.T) {
	payload := validPayload(false, 600)
	tokenBytes, _ := token.GenerateToken(payload)
	tokenBytes[5] ^= 0xFF // Flip one bit
	_, _, err := token.VerifyToken(tokenBytes)
	if err == nil {
		t.Errorf("Expected error for tampered token")
	}
}

// 🧪 Random bytes
func TestRandomBytesInput(t *testing.T) {
	random := []byte("garbage token data")
	_, _, err := token.VerifyToken(random)
	if err == nil {
		t.Errorf("Expected error for invalid token bytes")
	}
}

// 🧪 Large payload
func TestLargePayload(t *testing.T) {
	perms := make([]string, 0)
	for i := 0; i < 1000; i++ {
		perms = append(perms, "perm"+string(rune(i)))
	}
	payload := types.Payload{
		Sub:         "big-user",
		IssuedAt:    time.Now().Unix(),
		ExpiresAt:   time.Now().Add(time.Hour).Unix(),
		Nonce:       "big-nonce",
		SessionID:   "session-big",
		Permissions: perms,
	}
	_, err := token.GenerateToken(payload)
	if err != nil {
		t.Errorf("Expected large payload to be accepted, got error: %v", err)
	}
}
