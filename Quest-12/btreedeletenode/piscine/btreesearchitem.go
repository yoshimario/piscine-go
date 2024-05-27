package piscine

// BTreeSearchItem searches for a node with data equal to elem in the binary search tree.
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	current := root

	// Traverse the tree to find the node with the matching data.
	for current != nil {
		if elem == current.Data {
			return current
		}
		if elem < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// If the element is not found, return nil.
	return nil
}
