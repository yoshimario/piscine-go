package piscine

type NodeL2 struct {
	Data interface{}
	Next *NodeL2
}

type List2 struct {
	Head *NodeL2
	Tail *NodeL2
}

func ListPushFront(l *List2, data interface{}) {
	newNode := &NodeL2{Data: data}

	if l.Head == nil {
		// If the list is empty, both Head and Tail should point to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, add the new node at the beginning and update the Head
		newNode.Next = l.Head
		l.Head = newNode
	}
}
