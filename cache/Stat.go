package cache

import "log"

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) add(k string, v []byte) {
	s.Count++
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Count--
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
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
