package util

import "github.com/google/uuid"

func UUIDToBytes(id uuid.UUID) []byte {
	bytes := [16]byte(id)
	bytes2 := make([]byte, 16)
	copy(bytes2, bytes[:16])
	return bytes2
}
