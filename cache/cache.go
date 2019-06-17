package cache

import (
	"log"
)

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}

	if c == nil {
		panic("unknow type : " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
