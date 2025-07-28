package types

import "fmt"

type Payload struct {
	Sub         string   `json:"s" msgpack:"s"`
	IssuedAt    int64    `json:"i" msgpack:"i"`
	ExpiresAt   int64    `json:"e" msgpack:"e"`
	Nonce       string   `json:"n" msgpack:"n"`
	SessionID   string   `json:"r" msgpack:"r"`
	Permissions []string `json:"p" msgpack:"p"`
	Sliding     bool     `json:"sl" msgpack:"sl"`
}

func (p Payload) IsValid() bool {
	if p.Sub == "" || p.Nonce == "" || p.SessionID == "" || len(p.Permissions) == 0 || p.IssuedAt == 0 || p.ExpiresAt == 0 || p.ExpiresAt < p.IssuedAt {
		fmt.Println("Invalid payload:")
		return false
	}
	return true
}
