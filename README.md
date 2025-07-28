# 🔐 SwiftToken

**SwiftToken** is a blazing-fast, secure, and minimal token library for Go — designed as a lightweight and encrypted alternative to JWT.

---

## 🚀 Features

- 🔒 Uses [ChaCha20-Poly1305](https://pkg.go.dev/golang.org/x/crypto/chacha20poly1305) for encryption  
- ⚡ Up to **70x faster** than standard JWTs  
- 📦 Compact binary payloads via [MessagePack](https://msgpack.org/)  
- 🔁 Built-in sliding token support (auto refresh)  
- ✅ Fully tested against real-world edge cases  
- 📊 Benchmark comparisons with standard JWT

---

## 📦 Installation

Ensure your Go version is **1.20+**, then:

```bash
go get github.com/Aditya251610/swifttoken@v1.0.0
```

    💡 Tip: After installing, run this to ensure dependencies are properly resolved:
```bash
go mod tidy
```
---

## ⚙️ Setup

SwiftToken requires a 32-byte secret key.

### ✅ Option 1: Use `.env`

Create a `.env` file in your project root:

```env
SWIFTTOKEN_SECRET=12345678901234567890123456789012
```

### ✅ Option 2: Set via Code

```go
import "os"

os.Setenv("SWIFTTOKEN_SECRET", "12345678901234567890123456789012")
```

---

## 🧪 Usage

### 1️⃣ Generate a Token

```go
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
```

### 2️⃣ Verify a Token

```go
decoded, shouldRefresh, err := token.VerifyToken(tokenBytes)

if err != nil {
  // invalid, expired, or tampered token
}

if shouldRefresh {
  // token is nearing expiration, issue a new one
}
```

---

## 🔁 Sliding Tokens

Sliding tokens auto-refresh when close to expiration.

```go
payload.Sliding = true
```

On verification, SwiftToken will indicate if a new token should be issued.

---

## ✅ Testing

```bash
go test ./tests -v
```

All edge cases (tampering, expiration, etc.) are covered.

---

## ⚡ Benchmarks

```bash
go test -bench=. ./benchmarks -benchmem
```

| Function        | SwiftToken      | JWT (standard) |
|----------------|-----------------|----------------|
| Generate Token | ~669 ns/op      | ~4331 ns/op    |
| Verify Token   | ~82 ns/op       | ~5834 ns/op    |
| Allocations    | 6               | 49–70          |
| Payload Size   | Small (binary)  | Large (Base64) |

---

## 🛠️ Roadmap

- [ ] CLI tool for generating/verifying tokens  
- [ ] Token revocation support (via Redis/DB)  
- [ ] Token introspection endpoint  
- [ ] Plug-in based storage backends  
- [ ] SDK for frontend/mobile apps

---

## 🤝 Contributing

1. ⭐ Star the repo  
2. 🍴 Fork it  
3. Create your feature branch:
   ```bash
   git checkout -b feat/my-feature
   ```
4. Commit your changes:
   ```bash
   git commit -m "✨ add feature"
   ```
5. Push and open a PR ✅

---

## 📄 License

MIT © [Aditya Sharma](https://github.com/Aditya251610)

---

> SwiftToken — built for developers who care about **performance**, **security**, and **simplicity** 🔐⚡
