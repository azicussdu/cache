package cache

import "errors"

type Cache struct {
	cache map[string]any
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) error {
	if key == "" {
		return errors.New("key can not be empty")
	} else if value == nil {
		return errors.New("value can not be nil")
	}

	c.cache[key] = value
	return nil
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
