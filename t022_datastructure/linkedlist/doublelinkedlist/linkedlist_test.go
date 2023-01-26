package doublelinkedlist

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

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Nil(t, l.GetAt(3))
}

func TestLinkedList_PushFront(t *testing.T) {
	l := LinkedList[int]{}

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

	assert.Equal(t, 3, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	//assert.Equal(t, 1, l.GetAt(2).Value)
	//assert.Nil(t, l.GetAt(3))

}
