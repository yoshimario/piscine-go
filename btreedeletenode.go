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
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}


