package gache

import (
	"sync"

	"github.com/atsvetkov/gache/lru"
)

// ThreadSafeCache is a container for a specific cache implementation and a thread sync primitive
type ThreadSafeCache struct {
	cache Cache
	lock  sync.RWMutex
}

// NewLRUCache creates an thread-safe instance of a LRU (Least Recently Used) cache
func NewLRUCache(size int) (*ThreadSafeCache, error) {
	lruCache, err := lru.NewLRU(size)
	if err != nil {
		return nil, err
	}

	c := &ThreadSafeCache{cache: lruCache}

	return c, nil
}
