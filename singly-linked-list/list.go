package list

import "errors"

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

func (l *list) PushFront(v string) {
	l.len++
	l.root.next.InsertPrev(v)
}

func (l *list) PushBack(v string) {
	l.len++
	l.root.prev.InsertNext(v)
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

func (l *list) Get(index int) (string, error) {
	nodeAddr, err := l.get(index)
	if err != nil {
		return "", err
	}

	return nodeAddr.data, nil
}

//
//func (l *list) GetNodeByIndex(index int) (*node, error) {
//	return l.get(index)
//}
//
//func (l *list) Size() int {
//	return l.len
//}
//
//func (l *list) RemoveByIndex(index int) (string, error) {
//	nodeAddr, err := l.get(index)
//	if err != nil {
//		return "", err
//	}
//
//	l.len--
//
//	if nodeAddr.prev != nil {
//		nodeAddr.prev.next = nodeAddr.next
//	}
//
//	if nodeAddr.next != nil {
//		nodeAddr.next.prev = nodeAddr.prev
//	}
//
//	if nodeAddr == firstP {
//		if nodeAddr.next == nil {
//			firstP = nil
//		} else {
//
//		}
//	}
//
//	return nodeAddr.data, nil
//}
//
//func (l *list) Reverse() {
//
//}
