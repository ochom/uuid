package uuid

import (
	"encoding/binary"
	"encoding/hex"
	"time"
)

const (
	// ShortUID ...
	ShortUID = 12 // generate 24 characters long uuid
	// LongUID ...
	LongUID = 18 // generates 36 characters long uuid
)

// ObjectID ...
type ObjectID []byte

// UID ...
type UID struct {
	objectID ObjectID
}

// New ...
func New() ObjectID {
	b := make([]byte, ShortUID)
	str := generateRandomByte(4)
	copy(b, str)

	now := uint64(time.Now().Nanosecond())
	binary.BigEndian.PutUint64(b[4:], now)
	return b
}

// NewString ...
func NewString() string {
	id := New()
	return hex.EncodeToString(id)
}
