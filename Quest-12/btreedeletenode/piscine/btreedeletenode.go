package piscine

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Data   string
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode // Adding Parent pointer for easier deletion
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

// BTreeDeleteNode deletes the given node from the tree.
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Left == nil && node.Right == nil { // Case 1: Leaf node
		if node.Parent != nil {
			if node == node.Parent.Left {
				node.Parent.Left = nil
			} else {
				node.Parent.Right = nil
			}
		} else {
			root = nil
		}
	} else if node.Left == nil || node.Right == nil { // Case 2: Node has only one child
		var child *TreeNode
		if node.Left != nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.Parent != nil {
			if node == node.Parent.Left {
				node.Parent.Left = child
			} else {
				node.Parent.Right = child
			}
			child.Parent = node.Parent
		} else {
			root = child
			child.Parent = nil
		}
	} else { // Case 3: Node has two children
		successor := BTreeMin(node.Right)
		node.Data = successor.Data
		root = BTreeDeleteNode(root, successor)
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
