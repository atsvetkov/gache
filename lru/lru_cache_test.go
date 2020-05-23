package lru

import "testing"

func TestNewLRUCacheNeedsPositiveSize(t *testing.T) {
	_, err := NewLRU(0)
	if err == nil {
		t.Fail()
	}
}
