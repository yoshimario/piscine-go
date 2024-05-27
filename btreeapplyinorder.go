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

// BTreeApplyInorder applies the function f to each node in the tree in in-order.
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	// Recursively apply to the left subtree.
	BTreeApplyInorder(root.Left, f)
	// Apply the function to the current node's data.
	f(root.Data)
	// Recursively apply to the right subtree.
	BTreeApplyInorder(root.Right, f)
}
