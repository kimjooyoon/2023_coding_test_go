package doublelinkedlist

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l LinkedList[T]) Count() int {
	return l.count
}

func (l LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.count {
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

func (l *LinkedList[T]) PushBack(val T) {
	n := &Node[T]{Value: val}
	if l.root == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.count++
}

func (l *LinkedList[T]) PushFront(val T) {
	n := &Node[T]{Value: val}
	if l.root == nil {
		l.root, l.tail = n, n
		l.count = 1
		return
	}

	l.root.prev = n
	n.next = l.root
	l.root = n
	l.count++
}

func (l LinkedList[T]) isInclude(node *Node[T]) bool {
	for n := l.root; n != nil; n = n.next {
		if n == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isInclude(node) {
		return
	}
	n := &Node[T]{Value: val}

	nn := node.next
	n.next = nn
	node.next = n
	n.prev = node

	if nn != nil {
		nn.prev = n
	} else {
		l.tail = n
	}
	l.count++
}
