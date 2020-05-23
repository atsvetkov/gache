package gache

// Cache interface defines generic cache behavior
type Cache interface {
	Add(key, value interface{})
	Get(key interface{}) (value interface{}, ok bool)
	Contains(key interface{}) (ok bool)
	Remove(key interface{})
}
