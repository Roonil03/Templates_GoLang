package lru_cache

import (
	"container/list"
)

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*list.Element
	list     *list.List
}

type entry[K comparable, V any] struct {
	key K
	val V
}

func NewLRUCache[K comparable, V any](cap int) *LRUCache[K, V] {
	return &LRUCache[K, V]{capacity: cap, cache: make(map[K]*list.Element), list: list.New()}
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		return el.Value.(entry[K, V]).val, true
	}
	var zero V
	return zero, false
}

func (c *LRUCache[K, V]) Put(key K, val V) {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		el.Value = entry[K, V]{key, val}
		return
	}
	if c.list.Len() == c.capacity {
		el := c.list.Back()
		delete(c.cache, el.Value.(entry[K, V]).key)
		c.list.Remove(el)
	}
	en := entry[K, V]{key, val}
	el := c.list.PushFront(en)
	c.cache[key] = el
}
