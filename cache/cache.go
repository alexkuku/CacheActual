package cache

import "sync"

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

type inMemoryCache struct {
	c     map[string][]byte
	mutex sync.RWMutex
	Stat
}
