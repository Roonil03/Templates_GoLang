package ds

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type entry struct {
	key, val int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{capacity: cap, cache: make(map[int]*list.Element), list: list.New()}
}

func (c *LRUCache) Get(key int) (int, bool) {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		return el.Value.(entry).val, true
	}
	return 0, false
}

func (c *LRUCache) Put(key, val int) {
	if el, ok := c.cache[key]; ok {
		c.list.MoveToFront(el)
		el.Value = entry{key, val}
		return
	}
	if c.list.Len() == c.capacity {
		el := c.list.Back()
		delete(c.cache, el.Value.(entry).key)
		c.list.Remove(el)
	}
	en := entry{key, val}
	el := c.list.PushFront(en)
	c.cache[key] = el
}
