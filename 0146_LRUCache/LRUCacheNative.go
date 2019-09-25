package main

type LRUCacheNative struct {
	cache map[interface{}]*Node
	root *List
	size int
}

func ConstructorNative(size int) LRUCacheNative {
	instance := LRUCacheNative{
		cache: make(map[interface{}]*Node),
		root: NewList(),
		size: size,
	}
	return instance
}

func (this *LRUCacheNative) Put (key, value interface{})  {
	if node,ok := this.cache[key];ok {
		node.value = value
		this.root.moveToBack(node)
		return
	}
	if this.size == this.root.len {
		delNode := this.root.Front()
		this.root.remove(delNode)
		delete(this.cache, delNode.key)
	}
	node := &Node{
		key:   key,
		value: value,
		next:  nil,
		prev:  nil,
	}
	this.root.appendToBack(node)
	this.cache[key] = node
}

func (this *LRUCacheNative) Get(key interface{}) interface{} {
	if	node,ok := this.cache[key];ok {
		this.root.moveToBack(node)
		return node.value
	}
	return -1
}