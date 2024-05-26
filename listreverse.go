package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var prevNode *NodeL = nil
	currentNode := l.Head

	// Traverse the list
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	// Update the list's head and tail
	l.Tail, l.Head = l.Head, prevNode
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Tail == nil {
		// If the list is empty, both Head and Tail should point to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, add the new node at the end and update the Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
