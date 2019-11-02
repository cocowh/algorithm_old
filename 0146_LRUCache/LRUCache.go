package main

import "container/list"

type LRUCache struct {
	Size  int
	list  *list.List
	cache map[int]*list.Element
}

type Entry struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	instance := LRUCache{
		Size:  capacity,
		list:  list.New(),
		cache: make(map[int]*list.Element),
	}
	return instance
}

func (this *LRUCache) Get(key int) int {
	if this.cache == nil {
		return -1
	}
	if e, ok := this.cache[key]; ok {
		this.list.MoveToFront(e)
		return e.Value.(*Entry).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache == nil {
		this.cache = make(map[int]*list.Element)
		this.list = list.New()
	}
	if e, ok := this.cache[key]; ok {
		this.list.MoveToFront(e)
		e.Value.(*Entry).Value = value
		return
	}
	entry := &Entry{
		Key:   key,
		Value: value,
	}
	e := this.list.PushFront(entry)
	this.cache[key] = e
	if this.Size != 0 && this.list.Len() > this.Size {
		e := this.list.Back()
		if e != nil {
			this.list.Remove(e)
			delete(this.cache, e.Value.(*Entry).Key)
		}
	}
}
