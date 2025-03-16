package cache

import (
	"container/list"
	"sync"
)


type HexCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mu sync.Mutex
}

type Cacheable interface {
    GetID() string
}

func NewHexCache(capacity int) *HexCache {
	return &HexCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *HexCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		return el.Value, true
	}

	return nil, false
}

func (c *HexCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	
	if el, found := c.cache[key]; found {
		c.list.MoveToFront(el)
		el.Value = value
		return 
	}

	newEl := c.list.PushFront(value)
	c.cache[key] = newEl

	if c.list.Len() > c.capacity {	
		lastEl := c.list.Back()

		if lastEl != nil {
			delete(c.cache, lastEl.Value.(Cacheable).GetID())
			c.list.Remove(lastEl)
		}
	}
}

