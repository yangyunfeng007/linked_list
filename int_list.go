package linked_list

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

func NewInt() IntList {
	return newLinkedIntList()
}

type IntList interface {
	// 检查一个元素是否存在，如果存在则返回 true，否则返回 false
	Contains(value int) bool

	// 插入一个元素，如果此操作成功插入一个元素，则返回 true，否则返回 false
	Insert(value int) bool

	// 删除一个元素，如果此操作成功删除一个元素，则返回 true，否则返回 false
	Delete(value int) bool

	// 遍历此有序链表的所有元素，如果 f 返回 false，则停止遍历
	Range(f func(value int) bool)

	// 返回有序链表的元素个数
	Len() int
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func newIntNode(value int) *intNode {
	return &intNode{value: value}
}

type intNode struct {
	value int
	next  *intNode
	mu    sync.Mutex
	flags bitflag
}

// loadNext return `n.next`.
func (n *intNode) loadNext() *intNode {
	return (*intNode)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&n.next))))
}

// storeNext same with `n.next = node`.
func (n *intNode) storeNext(node *intNode) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&n.next)), unsafe.Pointer(node))
}

// lessthan same as n.value < value.
func (n *intNode) lessthan(value int) bool {
	return n.value < value
}

// equal same as n.value == value.
func (n *intNode) equal(value int) bool {
	return n.value == value
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

func newLinkedIntList() *linkedIntList {
	h := newIntNode(0)
	return &linkedIntList{
		header: h,
	}
}

type linkedIntList struct {
	header *intNode
	length int64
}

// Contains check if the value is in the list.
func (ll *linkedIntList) Contains(value int) bool {
	x := ll.header.loadNext()
	for x != nil && x.lessthan(value) {
		x = x.loadNext()
	}
	if x != nil && x.equal(value) {
		// Check x is not deleted in another thread.
		return !x.flags.Get(marked)
	}
	return false
}

// Insert insert the value into list, return true if this process insert the value into list,
// return false if this process can't insert this value, because another process has insert the same value.
//
// If the value is in the list but not fully linked, this process will wait until it is.
func (ll *linkedIntList) Insert(value int) bool {
FIND:
	// Step1: Find node a&b, and a.value < value < b.value.
	a := ll.header
	b := a.loadNext()
	for b != nil && b.lessthan(value) {
		a = b
		b = b.loadNext()
	}
	// Check the node exists and no duplicate value.
	if b != nil && b.equal(value) {
		return false
	}
	a.mu.Lock()
	// Step2: Check a is not deleted and a.next is stll b.
	if !a.flags.Get(marked) || a.loadNext() != b {
		a.mu.Unlock()
		goto FIND
	}
	// Step3: Insert x into a and b, a -> x -> b.
	x := newIntNode(value)
	x.storeNext(b)
	a.storeNext(x)
	a.mu.Unlock()
	atomic.AddInt64(&ll.length, 1)
	return true
}

// Delete a node from the list.
func (ll *linkedIntList) Delete(value int) bool {
FIND:
	// Step1: find a&b, and b.value == value.
	a := ll.header
	b := a.loadNext()
	for b != nil && b.lessthan(value) {
		a = b
		b = b.loadNext()
	}
	// Check b exists and no duplicate value.
	if b == nil || !b.equal(value) {
		return false
	}
	b.mu.Lock()
	if b.flags.Get(marked) {
		b.mu.Unlock()
		goto FIND
	}
	a.mu.Lock()
	// Step2: Check a.next is still b, and check a is not delete in another process.
	// If not, find a&b again, because b can be deleted in another thread.
	if a.next != b || a.flags.Get(marked) {
		a.mu.Unlock()
		b.mu.Unlock()
		goto FIND
	}
	// Step3: Delete b, and set b's marked to true(DELETED).
	// Set b's flag to DELETED.
	b.flags.SetTrue(marked)
	a.storeNext(b.loadNext())
	a.mu.Unlock()
	b.mu.Unlock()
	atomic.AddInt64(&ll.length, -1)
	return true
}

// Range calls f sequentially for each value present in the list.
// If f returns false, range stops the iteration.
func (ll *linkedIntList) Range(f func(value int) bool) {
	x := ll.header.loadNext()
	for x != nil {
		// Continue, if x is deleted in another thread.
		if x.flags.Get(marked) {
			x = x.loadNext()
			continue
		}
		if !f(x.value) {
			break
		}
		x = x.loadNext()
	}
}

// Len return the length of this list.
// Keep in sync with types_gen.go:lengthFunction
// Special case for code generation, Must in the tail of int_list.go.
func (ll *linkedIntList) Len() int {
	return int(atomic.LoadInt64(&ll.length))
}
