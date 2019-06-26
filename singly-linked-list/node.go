package list

type node struct {
	data string
	prev *node
	next *node
}

func NewNode(data string, prev *node, next *node) *node {
	return &node{
		data: data,
		prev: prev,
		next: next,
	}
}

func (nod *node) InsertPrev(newNode *node) *node {
	//var newNode node = *NewNode(v, nil, nod)

	if nod.prev != nil {
		newNode.prev = nod.prev
		nod.prev.next = newNode
	}

	nod.prev = newNode
}

func (nod *node) InsertNext(newNode *node) *node {
	//var newNode node = *NewNode(v, nod, nil)
	newNode.prev = nod
	newNode.next = nil

	if nod.next != nil {
		newNode.next = nod.next
		nod.next.prev = newNode
	}

	nod.next = newNode

	return &newNode
}
