package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	currentNode := l.Head

	for currentNode != nil {
		if comp(currentNode.Data, ref) {
			return &currentNode.Data
		}
		currentNode = currentNode.Next
	}

	return nil
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
