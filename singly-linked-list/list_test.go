package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestRemoveByIndex(t *testing.T) {
//	l := New()
//
//	l.PushBack("0")
//	l.PushBack("1")
//	l.PushBack("2")
//	l.PushBack("3")
//	l.PushBack("4")
//	l.PushBack("5")
//	l.PushBack("6")
//	l.PushBack("7")
//
//	l.RemoveByIndex(1)
//	l.RemoveByIndex(3)
//}

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

	// index out of range
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
