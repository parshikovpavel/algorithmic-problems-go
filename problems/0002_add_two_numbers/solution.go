package _0002_add_two_numbers

type Node struct {
	Val  int
	Next *Node
}

func AddTwoNumbers(node1 *Node, node2 *Node) *Node {
	var carry int

	current := &Node{}
	head := current

	for node1 != nil || node2 != nil || carry != 0 {
		val1 := 0
		val2 := 0
		if node1 != nil {
			val1 = node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			val2 = node2.Val
			node2 = node2.Next
		}
		val := val1 + val2 + carry

		carry = 0
		if val >= 10 {
			carry = 1
			val = val - 10
		}

		current.Next = &Node{
			val,
			nil,
		}
		current = current.Next

	}

	return head.Next
}
