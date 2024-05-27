package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeDeleteNode deletes a node from the binary tree
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	dataToMove := []string{}
	dataToMove = getTreeData(node.Right, &dataToMove)

	root = BTreeTransplant(root, node, node.Left)

	for _, st := range dataToMove {
		root = BTreeInsertData(root, st)
	}
	return root
}

// getTreeData recursively traverses the tree and collects data into a slice
func getTreeData(root *TreeNode, data *[]string) []string {
	if root == nil {
		return *data
	}
	*data = append(*data, root.Data)
	getTreeData(root.Left, data)
	getTreeData(root.Right, data)
	return *data
}
