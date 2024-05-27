package piscine

import "strconv"

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

// BTreeIsBinary checks if a tree is a binary search tree
func BTreeIsBinary(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

// isBST is a helper function to recursively check if a tree is a BST
func isBST(node *TreeNode, min *int, max *int) bool {
	if node == nil {
		return true
	}

	// Convert node.Data to an integer for comparison
	nodeValue, err := strconv.Atoi(node.Data)
	if err != nil {
		return false
	}

	// Check if the current node violates the min/max constraint
	if (min != nil && nodeValue <= *min) || (max != nil && nodeValue >= *max) {
		return false
	}

	// Recursively check the left and right subtrees
	return isBST(node.Left, min, &nodeValue) && isBST(node.Right, &nodeValue, max)
}

// BTreeInsertData inserts data into the binary search tree
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	rootValue, _ := strconv.Atoi(root.Data)
	dataValue, _ := strconv.Atoi(data)

	if dataValue < rootValue {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}

	return root
}
