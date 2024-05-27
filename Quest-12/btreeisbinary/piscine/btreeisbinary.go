package piscine

type TreeNode struct {
	Left, Right *TreeNode
	Data        string
}

// Insert data into the binary tree
func BTreeInsertData(root *TreeNode, data string) {
	if root == nil {
		return
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
}

// Helper function to check if the tree is a BST
func isBSTHelper(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}

	// Check if the current node's data is within the valid range
	if (min != nil && node.Data <= min.Data) || (max != nil && node.Data >= max.Data) {
		return false
	}

	// Recursively check the left and right subtrees with updated ranges
	return isBSTHelper(node.Left, min, node) && isBSTHelper(node.Right, node, max)
}

// Check if the binary tree is a BST
func BTreeIsBinary(root *TreeNode) bool {
	return isBSTHelper(root, nil, nil)
}
