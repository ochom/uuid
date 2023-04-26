package uuid

import "crypto/rand"

func generateRandomByte(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
	return b
}
