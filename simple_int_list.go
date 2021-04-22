package linked_list

import (
	"sync"
)

type SimpleIntList struct {
	head   *simpleIntNode
	length int64
	mu     sync.RWMutex
}

type simpleIntNode struct {
	value int
	next  *simpleIntNode
}

func newSimpleIntNode(value int) *simpleIntNode {
	return &simpleIntNode{value: value}
}

func NewSimpleInt() *SimpleIntList {
	return &SimpleIntList{head: newSimpleIntNode(0)}
}

func (l *SimpleIntList) Insert(value int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	a := l.head
	b := a.next
	for b != nil && b.value < value {
		a = b
		b = b.next
	}
	// Check if the node is exist.
	if b != nil && b.value == value {
		return false
	}
	x := newSimpleIntNode(value)
	x.next = b
	a.next = x
	l.length++
	return true
}

func (l *SimpleIntList) Delete(value int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	a := l.head
	b := a.next
	for b != nil && b.value < value {
		a = b
		b = b.next
	}
	// Check if b is not exists
	if b == nil || b.value != value {
		return false
	}
	a.next = b.next
	l.length--
	return true
}

func (l *SimpleIntList) Contains(value int) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	x := l.head.next
	for x != nil && x.value < value {
		x = x.next
	}
	if x == nil {
		return false
	}
	return x.value == value
}

func (l *SimpleIntList) Range(f func(value int) bool) {
	l.mu.RLock()
	x := l.head.next
	l.mu.RUnlock()
	for x != nil {
		if !f(x.value) {
			break
		}
		l.mu.RLock()
		x = x.next
		l.mu.RUnlock()
	}
}

func (l *SimpleIntList) Len() int {
	return int(l.length)
}
