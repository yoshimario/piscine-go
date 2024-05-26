package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	currentNode := l.Head

	for currentNode != nil {
		if cond(currentNode) {
			f(currentNode)
		}
		currentNode = currentNode.Next
	}
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int:
		return node.Data.(int) > 0
	case float32:
		return node.Data.(float32) > 0
	case float64:
		return node.Data.(float64) > 0
	case byte:
		return node.Data.(byte) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
