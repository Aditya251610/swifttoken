package benchmark

import (
	"testing"
	"time"

	"github.com/Aditya251610/swifttoken/token"
	"github.com/Aditya251610/swifttoken/types"
	"github.com/golang-jwt/jwt/v5"
)

var swiftPayload = types.Payload{
	Sub:         "user123",
	IssuedAt:    time.Now().Unix(),
	ExpiresAt:   time.Now().Add(15 * time.Minute).Unix(),
	Nonce:       "abc123",
	SessionID:   "sess1",
	Permissions: []string{"read", "write"},
	Sliding:     true,
}

var jwtSecret = []byte("myjwtsecretkey1234567890123456") // 32 bytes

func BenchmarkSwiftTokenGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = token.GenerateToken(swiftPayload)
	}
}

func BenchmarkSwiftTokenVerify(b *testing.B) {
	tok, _ := token.GenerateToken(swiftPayload)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = token.VerifyToken(tok)
	}
}

func BenchmarkJWTGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		claims := jwt.MapClaims{
			"sub":     swiftPayload.Sub,
			"iat":     swiftPayload.IssuedAt,
			"exp":     swiftPayload.ExpiresAt,
			"nonce":   swiftPayload.Nonce,
			"session": swiftPayload.SessionID,
			"perm":    swiftPayload.Permissions,
			"sliding": swiftPayload.Sliding,
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		_, _ = t.SignedString(jwtSecret)
	}
}

func BenchmarkJWTVerify(b *testing.B) {
	claims := jwt.MapClaims{
		"sub":     swiftPayload.Sub,
		"iat":     swiftPayload.IssuedAt,
		"exp":     swiftPayload.ExpiresAt,
		"nonce":   swiftPayload.Nonce,
		"session": swiftPayload.SessionID,
		"perm":    swiftPayload.Permissions,
		"sliding": swiftPayload.Sliding,
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := tok.SignedString(jwtSecret)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parsed, _ := jwt.Parse(ss, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		_, _ = parsed.Claims.(jwt.MapClaims)
	}
}
