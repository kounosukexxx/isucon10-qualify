package helper

import "sync"

type CacheMap struct {
	m map[interface{}]interface{}
	sync.RWMutex
}

func NewCacheMap() *CacheMap {
	return &CacheMap{
		m: make(map[interface{}]interface{}),
	}
}

func (c *CacheMap) Get(key interface{}) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.m[key]
	return value, ok
}

func (c *CacheMap) Set(key interface{}, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.m[key] = value
}

func (c *CacheMap) delete(key interface{}) {
	c.Lock()
	defer c.Unlock()
	delete(c.m, key)
}
