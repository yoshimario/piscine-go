package piscine

// TreeNode represents a node in a binary search tree.
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts new data into the binary search tree.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// Create a new node to be inserted.
	newNode := &TreeNode{Data: data}

	if root == nil {
		// If the root is nil, the new node becomes the root.
		return newNode
	}

	current := root
	var parent *TreeNode

	// Traverse the tree to find the correct position for the new node.
	for current != nil {
		parent = current
		if data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// Insert the new node as a child of the parent node.
	if data < parent.Data {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	// Set the parent of the new node.
	newNode.Parent = parent

	return root
}

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
