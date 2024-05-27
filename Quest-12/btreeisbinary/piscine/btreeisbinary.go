// File: btree.go

package piscine

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// BTreeIsBinary checks if a tree is a binary search tree
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}
	return true
}

// isBST is a helper function to recursively check if a tree is a BST
func isBST(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}

	// Check if the current node violates the min/max constraint
	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}

	// Recursively check the left and right subtrees
	return isBST(node.Left, min, &node.Data) && isBST(node.Right, &node.Data, max)
}

// BTreeInsertData inserts data into the binary search tree
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
