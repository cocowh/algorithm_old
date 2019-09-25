package main

type Node struct {
	key   interface{}
	value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	root Node
	len  int
}

func NewList() *List {
	l := new(List)
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Node {
	if l.len == 0{
		return nil
	}
	return l.root.next
}

func (l *List) insert (e,at *Node) *Node {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	l.len++
	return e
}

func (l *List) remove(e *Node)	*Node {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	l.len--
	return e
}

func (l *List) moveToBack(e *Node) {
	if l.root.prev == e {
		return
	}
	l.insert(l.remove(e),l.root.prev)
}

func (l *List) appendToBack(e *Node) {
	l.insert(e,l.root.prev)
}