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
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 1, l.Back().val)

	l.PushBack(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 2, l.Back().val)

	l.PushBack(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 3, l.Back().val)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	assert.Equal(t, 1, l.GetAt(0).val)
	assert.Equal(t, 2, l.GetAt(1).val)
	assert.Equal(t, 3, l.GetAt(2).val)
	assert.Nil(t, l.GetAt(3))
}

func TestLinkedList_PushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)
	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 1, l.Back().val)

	l.PushFront(2)
	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().val)
	assert.Equal(t, 1, l.Back().val)

	l.PushFront(3)
	assert.NotNil(t, l.root)
	assert.Equal(t, 3, l.Front().val)
	assert.Equal(t, 1, l.Back().val)

	assert.Equal(t, 3, l.Count())
	assert.Equal(t, 3, l.Count2())

	assert.Equal(t, 3, l.GetAt(0).val)
	assert.Equal(t, 2, l.GetAt(1).val)
	assert.Equal(t, 1, l.GetAt(2).val)
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
	assert.Equal(t, 4, l.GetAt(2).val)
	assert.Equal(t, 3, l.Back().val)

	tempNode := &Node[int]{val: 10}
	l.InsertAfter(tempNode, 200)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).val)
}

func TestLinkedList_InsertBefore(t *testing.T) {
	var l = LinkedList[int]{}

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.InsertBefore(l.GetAt(0), 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.Front().val)
}
