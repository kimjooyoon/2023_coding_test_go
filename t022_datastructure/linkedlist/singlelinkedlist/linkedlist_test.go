package singlelinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_PushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Nil(t, l.GetAt(3))
}

func TestLinkedList_PushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	assert.Equal(t, 3, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 1, l.GetAt(2).Value)
	assert.Nil(t, l.GetAt(3))
}

func TestLinkedList_InsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{Value: 10}
	l.InsertAfter(tempNode, 200)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).Value)
}

func TestLinkedList_InsertBefore(t *testing.T) {
	var l = LinkedList[int]{}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.InsertBefore(l.GetAt(1), 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{Value: 10}
	l.InsertBefore(tempNode, 200)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 2, l.GetAt(2).Value)
}

func TestLinkedList_InsertBeforeRoot(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.InsertBefore(l.GetAt(0), 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.GetAt(1).Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestLinkedList_PopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PopFront()
	assert.Equal(t, 2, l.Count2())

	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.PopFront()
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.PopFront()
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 0, l.Count2())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestLinkedList_Remove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(1))
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 1, l.Front().Value)

	assert.Equal(t, 3, l.Back().Value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.Remove(&Node[int]{
		Value: 100,
	})
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 0, l.Count2())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(2))
	assert.Equal(t, 2, l.Back().Value)
}

func TestLinkedList_Reverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
}

func TestLinkedList_Reverce2(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverce2()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
}
