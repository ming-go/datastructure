package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeInsertPrev(t *testing.T) {
	var nod *node = NewNode("1", nil, nil)

	var newNodeAddr1 *node = nod.InsertPrev("2")

	assert.Equal(t, nod, newNodeAddr1.next)
	assert.Equal(t, newNodeAddr1, nod.prev)
	assert.Equal(t, "2", nod.prev.data)

	var newNodeAddr2 *node = nod.InsertPrev("3")

	assert.Equal(t, newNodeAddr2, newNodeAddr1.next)
	assert.Equal(t, newNodeAddr1, nod.prev.prev)
}

func TestNodeInsertNext(t *testing.T) {
	var nod *node = NewNode("1", nil, nil)

	var newNodeAddr1 *node = nod.InsertNext("2")

	assert.Equal(t, nod, newNodeAddr1.prev)
	assert.Equal(t, newNodeAddr1, nod.next)
	assert.Equal(t, "2", nod.next.data)

	var newNodeAddr2 = nod.InsertNext("3")

	assert.Equal(t, newNodeAddr2, nod.next)
	assert.Equal(t, newNodeAddr1, nod.next.next)
}
