package LRU

import (
	"container/list"
)

type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

type Cache struct {
	MaxCount   int
	OnEnvicted func(key Key, value interface{})
	list       *list.List
	cache      map[interface{}]*list.Element
}

func NewCache(cap int) *Cache {
	return &Cache{
		MaxCount: cap,
		list:     list.New(),
		cache:    make(map[interface{}]*list.Element),
	}
}

func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.list = list.New()
	}

	if ele, hit := c.cache[key]; hit {
		c.list.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	ele := c.list.PushFront(&entry{key, value})
	c.cache[key] = ele

	//判断是否超出最大值
	if c.MaxCount > 0 && c.Len() > c.MaxCount {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}

	if ele, hit := c.cache[key]; hit {
		c.list.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.list.Back()
	c.removeElement(ele)
}

func (c *Cache) removeElement(element *list.Element) {
	c.list.Remove(element)
	kv := element.Value.(*entry)
	delete(c.cache, kv.key)
	if c.OnEnvicted != nil {
		c.OnEnvicted(kv.key, kv.value)
	}
}

func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.list.Len()
}
