package lrumap

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

type item struct {
	value interface{}
}

type LruMap struct {
	lru *lru.Cache[any, item]
}

// New creates and returns a new instance of a TTLMap.
// It takes in a `maxTTL` integer that represents the maximum time-to-live
// in milliseconds for each key in the TTLMap.
// It returns a pointer to the newly created TTLMap.
func New(capacity int) (m *LruMap) {
	var c int
	if capacity < 1000 {
		c = 1000
	} else {
		c = capacity
	}

	lru, err := lru.New[any, item](c)
	if err != nil {
		panic("lru error")
	}

	m = &LruMap{lru}
	return
}

func (m *LruMap) Get(key interface{}) (interface{}, bool) {
	item, ok := m.lru.Get(key)
	if !ok {
		return nil, false
	}
	return item.value, true
}

func (m *LruMap) Delete(key interface{}) {
	m.lru.Remove(key)
}

func (m *LruMap) Set(key interface{}, value interface{}) {
	m.lru.Add(key, item{value: value})
}

func (m *LruMap) Purge() {
	m.lru.Purge()
}
