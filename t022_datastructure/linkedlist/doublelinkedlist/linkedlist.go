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
