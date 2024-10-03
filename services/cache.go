package services

import (
	"sync"
	"time"
)

type CacheItem struct {
	Data      interface{}
	ExpiredAt time.Time
}

type CacheService struct {
	data sync.Map
}

func NewCacheService() *CacheService {
	return &CacheService{
		data: sync.Map{},
	}
}

func (c *CacheService) Set(key string, value interface{}, duration time.Duration) {
	c.data.Store(key, CacheItem{
		Data:      value,
		ExpiredAt: time.Now().Add(duration),
	})
}

func (c *CacheService) Get(key string) (interface{}, bool) {
	item, ok := c.data.Load(key)
	if !ok {
		return nil, false
	}
	cacheItem := item.(CacheItem)
	if time.Now().After(cacheItem.ExpiredAt) {
		c.data.Delete(key)
		return nil, false
	}
	return cacheItem.Data, true
}
