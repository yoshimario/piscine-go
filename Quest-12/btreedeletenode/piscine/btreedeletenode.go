package piscine

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
