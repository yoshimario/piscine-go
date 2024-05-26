package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Initialize a dummy node to facilitate merging
	dummy := &NodeI{}
	current := dummy

	// Iterate through both lists until one of them becomes nil
	for n1 != nil && n2 != nil {
		// Compare the data of nodes from both lists
		if n1.Data <= n2.Data {
			// Append the node from n1 to the merged list
			current.Next = n1
			n1 = n1.Next
		} else {
			// Append the node from n2 to the merged list
			current.Next = n2
			n2 = n2.Next
		}
		// Move the current pointer forward
		current = current.Next
	}

	// Append the remaining nodes from n1 or n2 to the merged list
	if n1 != nil {
		current.Next = n1
	} else {
		current.Next = n2
	}

	// Return the merged list, skipping the dummy node
	return dummy.Next
}
