package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		ll:       list.List{},
		items:    make(map[K]*list.Element),
	}
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	element, ok := c.items[key]
	if !ok {
		return value, false
	}
	c.ll.MoveToFront(element)
	value = element.Value.(*entry[K, V]).value
	return value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}
	_, exists := c.items[key]

	if !exists {
		if len(c.items) >= c.capacity {
			lastElement := c.ll.Back()
			c.ll.Remove(lastElement)
			delete(c.items, lastElement.Value.(*entry[K, V]).key)
		}
		record := &entry[K, V]{key: key, value: value}	//в этой переменной храню указатель на структуру entry
		element := c.ll.PushFront(record)				//в этой переменной уже указатель на элемент в списке
		c.items[key] = element							//по ключу получаю указатель
	} else {
		element := c.items[key]
		element.Value.(*entry[K, V]).value = value
	}
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
