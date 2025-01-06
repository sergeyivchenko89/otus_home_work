package hw04lrucache

type Key string

type KeyValuePair struct {
	key   Key
	value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	list

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Clear() {
	c.queue.Front().Next = nil
	c.queue.Back().Prev = nil
	c.items = nil
}

func (c lruCache) Get(key Key) (interface{}, bool) {

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	keyValue := item.Value.(KeyValuePair)
	c.queue.MoveToFront(item)
	return keyValue.value, true
}

func (c *lruCache) Set(key Key, value interface{}) bool {

	item, ok := c.items[key]
	v := KeyValuePair{key, value}
	if ok {
		item.Value = v
		c.queue.PushFront(item)
		return true
	}

	if c.queue.Len() == c.capacity {
		lastNode := c.queue.Back()
		lastNodeValue := lastNode.Value.(KeyValuePair)
		lastKey := lastNodeValue.key
		delete(c.items, lastKey)
		c.queue.Remove(lastNode)
	}

	c.items[key] = c.queue.PushFront(v)

	return false
}
