package piscine

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// BTreeInsertData inserts a new node with the given data into the binary tree.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

// BTreeMin returns the node with the minimum value in the tree.
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}
