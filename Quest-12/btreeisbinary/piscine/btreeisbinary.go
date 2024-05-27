package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Data < root.Left.Data {
		return false
	}
	if root.Right != nil && root.Data > root.Right.Data {
		return false
	}
	left := BTreeIsBinary(root.Left)
	right := BTreeIsBinary(root.Right)

	if !left || !right {
		return false
	}
	return true
}
