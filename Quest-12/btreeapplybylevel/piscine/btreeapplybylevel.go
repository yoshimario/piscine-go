package piscine

import (
	"fmt"
)

type TreeNode struct {
	Left, Right *TreeNode
	Data        string
}

// Insert data into the binary tree
func BTreeInsertData(root *TreeNode, data string) {
	if root == nil {
		return
	}

	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
}

// Apply the function f to each node of the tree level by level
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Create a queue to store nodes at each level
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Get the first node in the queue
		node := queue[0]
		queue = queue[1:]

		// Apply the function f to the current node's data
		f(node.Data)

		// Add the left and right children to the queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

// Main function to test BTreeApplyByLevel
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	BTreeApplyByLevel(root, fmt.Println)
}
