package piscine

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
