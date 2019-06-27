package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	errors "golang.org/x/xerrors"
)

func TestInsertBefore(t *testing.T) {
	l := New()

	newNode := l.insertBefore(&node{data: "1"}, &l.root)
	assert.Equal(t, "1", l.root.next.data)
	assert.Equal(t, &l.root, newNode.prev)
	assert.Equal(t, &l.root, newNode.next)
	assert.Equal(t, newNode, l.root.next)
	assert.Equal(t, newNode, l.root.prev)
	assert.Equal(t, l, newNode.list)
	assert.Equal(t, 1, l.len)

	newNode = l.insertBefore(&node{data: "2"}, newNode)
	newNode = l.insertBefore(&node{data: "3"}, newNode)
	newNode = l.insertBefore(&node{data: "4"}, newNode)
	newNode = l.insertBefore(&node{data: "5"}, newNode)
	newNode = l.insertBefore(&node{data: "6"}, newNode)
	newNode = l.insertBefore(&node{data: "8"}, newNode)

	assert.Equal(t, "8", l.root.next.next.next.next.next.next.next.data)

	newNode = l.insertBefore(&node{data: "7"}, newNode.prev)

	assert.Equal(t, "6", l.root.next.next.next.next.next.next.data)
	assert.Equal(t, "7", l.root.next.next.next.next.next.next.next.data)
	assert.Equal(t, "8", l.root.next.next.next.next.next.next.next.next.data)
}

func TestGet(t *testing.T) {
	l := New()

	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")
	l.PushBack("4")
	l.PushBack("5")
	l.PushBack("6")
	l.PushBack("7")

	v, err := l.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "1", v)
	v, err = l.Get(6)
	assert.Nil(t, err)
	assert.Equal(t, "7", v)

	// index out of range
	v, err = l.Get(8)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(ErrIndexOutOfRange, err))
}

func TestPushBack(t *testing.T) {
	l := New()

	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")
	l.PushBack("4")
	l.PushBack("5")
	l.PushBack("6")
	l.PushBack("7")

	v, err := l.Get(2)
	assert.Nil(t, err)
	assert.Equal(t, "3", v)

	assert.Equal(t, 7, l.len)
}

func TestPushFront(t *testing.T) {
	l := New()

	l.PushFront("1")
	l.PushFront("2")
	l.PushFront("3")
	l.PushFront("4")
	l.PushFront("5")
	l.PushFront("6")
	l.PushFront("7")

	v, err := l.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, "7", v)

	assert.Equal(t, 7, l.len)
}

func TestRemove(t *testing.T) {
	l := New()

	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")
	l.PushBack("4")
	l.PushBack("5")
	l.PushBack("6")
	l.PushBack("7")

	nod, err := l.get(2)
	assert.Nil(t, err)
	v := l.remove(nod)
	assert.Nil(t, err)
	assert.Equal(t, "3", v)
}

func TestRemoveByIndex(t *testing.T) {
	l := New()

	l.PushBack("1")
	l.PushBack("2")
	l.PushBack("3")
	l.PushBack("4")
	l.PushBack("5")
	l.PushBack("6")
	l.PushBack("7")

	v, err := l.RemoveByIndex(3)
	assert.Nil(t, err)
	assert.Equal(t, 6, l.len)
	assert.Equal(t, "4", v)

	vv, err := l.get(2)
	assert.Equal(t, "3", vv.data)

	vv, err = l.get(3)
	assert.Equal(t, "5", vv.data)

	v, err = l.RemoveByIndex(3)
	assert.Nil(t, err)
	assert.Equal(t, 5, l.len)
	assert.Equal(t, "5", v)
}
