package cache

type Cache interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
}