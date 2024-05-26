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

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove nodes from the head of the list if they match data_ref
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If the list becomes empty
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Remove nodes from the rest of the list
	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Data == data_ref {
			prev.Next = curr.Next
			if curr.Next == nil {
				l.Tail = prev
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}
}
