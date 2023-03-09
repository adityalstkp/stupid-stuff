package stupidstuff

import (
	"fmt"
	"sync"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

type Cacher interface {
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	val, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", key)
	}

	return val, nil
}

func (c *Cache) Set(key string, value []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = value
	return nil
}

func (c *Cache) Delete(key string) error {
	c.lock.Lock()
	c.lock.RLock()

	defer c.lock.RUnlock()
	defer c.lock.Unlock()

	delete(c.data, key)

	return nil

}
