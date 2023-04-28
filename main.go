package uuid

import (
	"encoding/binary"
	"encoding/hex"
	"sync/atomic"
	"time"
)

var objectIDCounter = readRandomUint32()

// ObjectID ...
type ObjectID [12]byte

// New ...
func New() ObjectID {
	var b [12]byte

	// Convert the time to a byte array
	binary.BigEndian.PutUint32(b[0:4], uint32(time.Now().Unix()))

	// Generate 5 random bytes
	unique := generateRandomByte()
	copy(b[4:9], unique[:])

	now := uint64(time.Now().UnixNano())
	binary.BigEndian.PutUint64(b[4:], now)

	// Generate 3 random bytes with counter
	putUint32(b[9:12], atomic.AddUint32(&objectIDCounter, 1))

	return b
}

// NewString ...
func NewString() string {
	id := New()
	return hex.EncodeToString(id[:])
}
