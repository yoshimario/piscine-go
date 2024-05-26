package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

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

func ListSize(l *List) int {
	count := 0
	for node := l.Head; node != nil; node = node.Next {
		count++
	}
	return count
}
