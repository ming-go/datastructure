/*
	Ref:
		[1]: https://golang.org/pkg/container/list
*/

package list

import (
	errors "golang.org/x/xerrors"
)

var (
	ErrIndexOutOfRange = errors.New("Index out of range")
)

type list struct {
	root node
	len  int
}

func New() *list {
	return new(list).Init()
}

func (l *list) Init() *list {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *list) insertBefore(newNode, at *node) *node {
	newNode.next = at.next
	at.next.prev = newNode
	at.next = newNode
	newNode.prev = at
	newNode.list = l
	l.len++
	return newNode
}

func (l *list) insertValueBefore(v interface{}, at *node) *node {
	return l.insertBefore(&node{data: v}, at)
}

func (l *list) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *list) Front() *node {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *list) Back() *node {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

func (l *list) PushFront(v interface{}) {
	l.lazyInit()
	l.insertValueBefore(v, &l.root)
}

func (l *list) PushBack(v interface{}) {
	l.lazyInit()
	l.insertValueBefore(v, l.root.prev)
}

//
func (l *list) get(index int) (*node, error) {
	if index > (l.len - 1) {
		return nil, ErrIndexOutOfRange
	}

	var nodeAddr *node = l.root.next

	for index > 0 {
		nodeAddr = nodeAddr.next
		index--
	}

	return nodeAddr, nil
}

func (l *list) Get(index int) (interface{}, error) {
	nodeAddr, err := l.get(index)
	if err != nil {
		return "", err
	}

	return nodeAddr.data, nil
}

func (l *list) remove(nod *node) interface{} {
	nod.prev.next = nod.next
	nod.next.prev = nod.prev
	nod.prev = nil // avoid memory leaks
	nod.next = nil // avoid memory leaks
	nod.list = nil // avoid memory leaks
	l.len--

	return nod.data
}

func (l *list) RemoveByIndex(index int) (interface{}, error) {
	nodeAddr, err := l.get(index)
	if err != nil {
		return -1, err
	}

	return l.remove(nodeAddr), nil
}

func (l *list) Sort() {

}

func (l *list) Reverse() {

}

func (l *list) Search() {

}

func (l *list) DeDuplicate() {

}
