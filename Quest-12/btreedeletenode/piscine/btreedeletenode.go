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

func getTreeData(node *TreeNode, dataToMove *[]string) []string {
	if node != nil {
		*dataToMove = append(*dataToMove, node.Data)
		getTreeData(node.Left, dataToMove)
		getTreeData(node.Right, dataToMove)
	}
	return *dataToMove
}
