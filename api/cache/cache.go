package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	Store *cache.Cache
)


type Cacher interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Flush()
}

const (
	// For use with functions that take an expiration time.
	RememberForever time.Duration = -1
)

// Init set-ups go-cache with defaults
func Init() {
	Store = cache.New(5*time.Minute, 10*time.Minute)
}
