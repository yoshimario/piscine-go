package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	// If the list is empty or the new node's data is smaller than the head's data,
	// insert the new node at the beginning and update the head
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	// Traverse the list to find the appropriate position to insert the new node
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert the new node after the current node
	newNode.Next = current.Next
	current.Next = newNode

	return l
}
