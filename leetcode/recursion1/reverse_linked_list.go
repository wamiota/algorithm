package recursion1

import (
	"algorithm/leetcode/structure"
)

func reverseList(head *structure.ListNode) *structure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return reverseListRecursive(head, nil, head.Next)
}

func reverseListRecursive(node *structure.ListNode, prev *structure.ListNode, next *structure.ListNode) *structure.ListNode {
	if node.Next == nil {
		node.Next = prev
		return node
	}

	node.Next = prev
	prev = node
	node = next
	return reverseListRecursive(node, prev, node.Next)
}

func reverseListIterate(head *structure.ListNode) *structure.ListNode {
	var prev *structure.ListNode
	var next *structure.ListNode

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
