package singlelinkedlist

type Node[T any] struct {
	next *Node[T]
	val  T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	n := &Node[T]{val: value}
	l.count++
	if l.root == nil {
		l.root = n
		l.tail = n
		return
	}
	l.tail.next = n
	l.tail = n
}

func (l *LinkedList[T]) PushFront(value T) {
	n := &Node[T]{val: value}
	l.count++
	if l.root == nil {
		l.root = n
		l.tail = n
		return
	}
	n.next = l.root
	l.root = n
}

func (l LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l LinkedList[T]) Count() int {
	n := l.root
	cnt := 0
	for ; n != nil; n = n.next {
		cnt++
	}
	return cnt
}

func (l LinkedList[T]) Count2() int {
	return l.count
}

func (l LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count2() {
		return nil
	}
	i := 0
	for n := l.root; n != nil; n = n.next {
		if i == idx {
			return n
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isCluded(node) {
		return
	}
	n := &Node[T]{val: value}

	n.next, node.next = node.next, n
	l.count++
}

func (l LinkedList[T]) isCluded(node *Node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if n == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if l.root == nil {
		l.PushFront(value)
		return
	}
	pre := l.findPrevNode(node)
	if pre == nil {
		return
	}

	n := &Node[T]{val: value}

	pre.next, n.next = n, node
	l.count++
}

func (l LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	var pre = l.root
	for ; pre != nil; pre = pre.next {
		if pre.next == node {
			return pre
		}
	}
	return nil
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	n := l.root
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--

	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}
	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}
	prev.next, node.next = node.next, nil
	if node == l.tail {
		l.tail = prev
	}
	l.count--
}

func (l *LinkedList[T]) Reverse() {
	l2 := &LinkedList[T]{}
	for l.root != nil {
		l2.PushFront(l.PopFront().val)
	}
	l.root, l.tail = l2.root, l2.tail
}

func (l *LinkedList[T]) Reverce2() {
	if l.root == nil {
		return
	}

	c := l.root
	n := c.next
	l.root.next = nil
	for n != nil {
		nn := n.next
		n.next = c
		c, n = n, nn
	}

	l.root.next = nil
}
