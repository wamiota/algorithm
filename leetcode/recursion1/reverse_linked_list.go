package recursion1

import (
	"algorithm/leetcode/ds"
)

func reverseList(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	return reverseListRecursive(head, nil, head.Next)
}

func reverseListRecursive(node *ds.ListNode, prev *ds.ListNode, next *ds.ListNode) *ds.ListNode {
	if node.Next == nil {
		node.Next = prev
		return node
	}

	node.Next = prev
	prev = node
	node = next
	return reverseListRecursive(node, prev, node.Next)
}

func reverseListIterate(head *ds.ListNode) *ds.ListNode {
	var prev *ds.ListNode
	var next *ds.ListNode

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
