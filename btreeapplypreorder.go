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

// Apply the function f to each node of the tree using preorder traversal
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
