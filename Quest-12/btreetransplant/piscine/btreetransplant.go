package piscine

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode // Adding Parent pointer for easier transplant
}

// BTreeInsertData inserts a new node with the given data into the binary tree.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}

// BTreeSearchItem searches for a node with the given data in the binary tree.
func BTreeSearchItem(root *TreeNode, data string) *TreeNode {
	if root == nil || root.Data == data {
		return root
	}
	if data < root.Data {
		return BTreeSearchItem(root.Left, data)
	}
	return BTreeSearchItem(root.Right, data)
}

// BTreeApplyInorder applies a function to each node of the binary tree in order.
func BTreeApplyInorder(root *TreeNode, fn func(string)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, fn)
	fn(root.Data)
	BTreeApplyInorder(root.Right, fn)
}

// BTreeTransplant replaces the subtree rooted at node with the subtree rooted at rplc.
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
