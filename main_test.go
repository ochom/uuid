package uuid

import (
	"testing"
)

func TestNew(t *testing.T) {
	hash := map[string]bool{}
	for i := 0; i < 1000000; i++ {
		id := NewString()
		if _, ok := hash[id]; ok {
			t.Errorf("object id %s repeated", id)
			break
		}

		hash[id] = true
	}
}
