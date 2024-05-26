package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListForEach(l *List, f func(*NodeL)) {
	currentNode := l.Head

	for currentNode != nil {
		f(currentNode)
		currentNode = currentNode.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Tail == nil {
		// If the list is empty, both Head and Tail should point to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// Otherwise, add the new node at the end and update the Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
