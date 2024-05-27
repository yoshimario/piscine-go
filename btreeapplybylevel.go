package piscine

import "fmt"

// BTreeApplyByLevel applies the function f to each node of the tree
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Create a queue to store nodes for level-order traversal
	queue := []*TreeNode{root}

	// Iterate through the queue until it's empty
	for len(queue) > 0 {
		// Pop the front node from the queue
		node := queue[0]
		queue = queue[1:]

		// Apply the function f to the node's data
		if _, err := f(node.Data); err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Add the left and right children of the current node to the queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
