package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	dataToMove := []string{}
	dataToMove = getTreeData(node.Right, &dataToMove)

	root = BTreeTransplant(root, node, node.Left)

	for _, st := range dataToMove {
		BTreeInsertData(root, st)
	}
	return root
}

func getTreeData(node *TreeNode, data *[]string) []string {
	if root == nil {
		return *data
	}
	*data = append(*data, root.Data)
	getTreeData(root.Left, data)
	getTreeData(root.Right, data)
	return *data
}
