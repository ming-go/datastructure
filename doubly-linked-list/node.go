package list

type node struct {
	data interface{}
	prev *node
	next *node
	list *list
}

func NewNode(data interface{}, prev *node, next *node) *node {
	return &node{
		data: data,
		prev: prev,
		next: next,
	}
}

//func (nod *node) InsertPrev(newNode *node) *node {
//	if nod.prev == nil {
//		nod.prev = nod
//	}
//
//	return nod.prev.InsertNext(newNode)
//}
//
//func (nod *node) InsertNext(newNode *node) *node {
//	newNode.prev = nod
//	newNode.next = nil
//
//	if nod.next != nil {
//		newNode.next = nod.next
//		nod.next.prev = newNode
//	}
//
//	nod.next = newNode
//
//	return newNode
//}
