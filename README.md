🔐 SwiftToken

SwiftToken is a blazing-fast, secure, and minimal token library for Go — designed as a lightweight and encrypted alternative to JWT.

    🔒 Uses ChaCha20-Poly1305 for encryption

    ⚡ Up to 70x faster than standard JWTs

    📦 Compact payloads using MessagePack

    🔁 Built-in sliding token support

    ✅ Fully tested with real-world edge cases

    📊 Benchmark comparisons with JWT

📦 Installation

Make sure your Go version is 1.20+ and then:

go get github.com/Aditya251610/swifttoken@v1.0.0

⚙️ Setup

SwiftToken requires a 32-byte secret key.

✅ Option 1: Use .env

Create a .env file in your project root:

SWIFTTOKEN_SECRET=12345678901234567890123456789012

✅ Option 2: Set via code

import "os"

os.Setenv("SWIFTTOKEN_SECRET", "12345678901234567890123456789012")

🚀 Usage
1️⃣ Generate a Token

import (
  "time"
  "github.com/Aditya251610/swifttoken/token"
  "github.com/Aditya251610/swifttoken/types"
)

func main() {
  payload := types.Payload{
    Sub:         "user123",
    IssuedAt:    time.Now().Unix(),
    ExpiresAt:   time.Now().Add(30 * time.Minute).Unix(),
    SessionID:   "sess-xyz",
    Permissions: []string{"read", "write"},
    Sliding:     true,
    Nonce:       "1234abcd",
  }

  tokenBytes, err := token.GenerateToken(payload)
  if err != nil {
    panic(err)
  }
}

2️⃣ Verify a Token

decoded, shouldRefresh, err := token.VerifyToken(tokenBytes)

if err != nil {
  // invalid, expired or tampered token
}

if shouldRefresh {
  // token is nearing expiration, issue a new one
}

🔁 Sliding Tokens

Sliding tokens auto-trigger refresh logic if they're nearing expiration.

payload.Sliding = true

On verification, SwiftToken will tell you if the token should be refreshed.
🧪 Testing

go test ./tests -v

All edge cases and attacks (tampering, expired tokens, etc.) are covered.
⚡ Benchmarks

Run the benchmarks:

go test -bench=. ./benchmarks -benchmem

Function
	

SwiftToken
	

JWT (standard)

Generate Token
	

~669 ns/op
	

~4331 ns/op

Verify Token
	

~82 ns/op
	

~5834 ns/op

Allocations
	

6
	

49–70

Payload Size
	

Small (binary)
	

Large (Base64)
🛠️ Roadmap

    [ ] CLI tool for generating/verifying tokens

    [ ] Revocation support (via Redis or DB)

    [ ] Token introspection endpoint

    [ ] Plug-in based storage backend

    [ ] SDK for frontend/mobile apps

🤝 Contributing

    Star the repo ⭐

    Fork it 🍴

    Create your feature branch git checkout -b feat/my-feature

    Commit changes git commit -m "✨ add feature"

    Push and open a PR ✅

📄 License

MIT © Aditya Sharma