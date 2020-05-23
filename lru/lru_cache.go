package lru

import (
	"errors"
)

type lru struct {
	size  int
	items map[interface{}]interface{}
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewLRU(size int) (*lru, error) {
	if size <= 0 {
		return nil, errors.New("Cache size must be greater than zero")
	}

	cache := lru{
		size:  size,
		items: make(map[interface{}]interface{})}

	return &cache, nil
}

func (c *lru) Add(key, value interface{}) {
	c.items[key] = value
}

func (c *lru) Get(key interface{}) (interface{}, bool) {
	if value, ok := c.items[key]; ok {
		return value, true
	}

	return nil, false
}

func (c *lru) Contains(key interface{}) bool {
	_, ok := c.items[key]

	return ok
}

func (c *lru) Remove(key interface{}) {
	delete(c.items, key)
}
