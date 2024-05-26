package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	// Split the list into two halves
	mid := getMiddle(l)
	nextToMid := mid.Next
	mid.Next = nil

	// Recursively sort the two halves
	left := ListSort(l)
	right := ListSort(nextToMid)

	// Merge the sorted halves
	sortedList := merge(left, right)
	return sortedList
}

func getMiddle(head *NodeI) *NodeI {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *NodeI) *NodeI {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var result *NodeI

	if left.Data <= right.Data {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}

	return result
}
