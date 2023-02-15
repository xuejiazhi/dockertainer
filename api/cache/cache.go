package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	NormalCache   *cache.Cache
	NormalTimeOut = int64(30)
)

func InitCache() {
	NormalCache = cache.New(5*time.Minute, 10*time.Minute)
}

// Set  value
func Set(key string, value interface{}, ttl int64) {
	if ttl == 0 {
		ttl = NormalTimeOut
	}
	NormalCache.Set(key, value, time.Duration(ttl)*time.Second)
}

// Get cache value
func Get(key string) (interface{}, bool) {
	data, ok := NormalCache.Get(key)
	return data, ok
}
