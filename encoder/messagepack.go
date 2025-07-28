package encoder

import (
	"github.com/Aditya251610/swifttoken/types"
	"github.com/vmihailenco/msgpack/v5"
)

func EncodePayload(payload types.Payload) ([]byte, error) {
	return msgpack.Marshal(payload)
}

func DecodePayload(data []byte) (types.Payload, error) {
	var p types.Payload
	err := msgpack.Unmarshal(data, &p)
	return p, err
}
