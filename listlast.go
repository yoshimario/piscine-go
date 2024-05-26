package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		// If the list is empty, return nil
		return nil
	}
	return l.Tail.Data
}
