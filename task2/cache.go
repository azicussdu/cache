package task2

import (
	"context"
	"errors"
	"sync"
	"time"
)

// added Mutex field for synchronization issues
type Cache struct {
	cache map[string]any
	mu    *sync.Mutex
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]any),
		mu:    new(sync.Mutex),
	}
}

func (c *Cache) Set(key string, value any, tm time.Duration) error {
	if key == "" {
		return errors.New("key can not be empty")
	} else if value == nil {
		return errors.New("value can not be nil")
	}

	c.cache[key] = value
	go c.RemoveAfter(key, tm)

	return nil
}

func (c *Cache) RemoveAfter(key string, tm time.Duration) {
	ctxT, cancel := context.WithTimeout(context.Background(), tm)
	defer cancel()

	for {
		select {
		case <-ctxT.Done():
			c.mu.Lock()
			delete(c.cache, key)
			c.mu.Lock()

			return
		}
	}
}

func (c *Cache) Get(key string) (any, error) {
	if key == "" {
		return nil, errors.New("key can not be empty")
	}

	val, ok := c.cache[key]
	if !ok {
		return nil, errors.New("there is no value with this key")
	}
	return val, nil
}

func (c *Cache) Delete(key string) error {
	if key == "" {
		return errors.New("key can not be empty")
	}

	_, ok := c.cache[key]
	if !ok {
		return errors.New("there is no value with this key")
	}

	delete(c.cache, key)
	return nil
}
