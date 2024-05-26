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
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		// If the first list is empty, simply point l1's Head and Tail to l2's Head and Tail
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else if l2.Head != nil {
		// If the first list is not empty and the second list is not empty,
		// append the second list to the end of the first list
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
	// Clear l2
	l2.Head = nil
	l2.Tail = nil
}
