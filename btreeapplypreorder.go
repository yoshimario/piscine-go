package piscine

type TreeNode struct {
	Left, Right *TreeNode
	Data        string
}

// BTreeInsertData inserts data into the binary tree.
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

// BTreeApplyPreorder applies the function f to each node of the tree using preorder traversal.
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Apply the function f to the current node's data
	f(root.Data)

	// Recursively apply the function to the left and right subtrees
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
