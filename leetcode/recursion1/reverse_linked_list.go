package recursion1

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return reverseListRecursive(head, nil, head.Next)
}

func reverseListRecursive(node *ListNode, prev *ListNode, next *ListNode) *ListNode {
	if node.Next == nil {
		node.Next = prev
		return node
	}

	node.Next = prev
	prev = node
	node = next
	return reverseListRecursive(node, prev, node.Next)
}

func reverseListIterate(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode

	for {
		if head.Next == nil {
			head.Next = prev
			return head
		}
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
}
