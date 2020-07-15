package recursion1

import (
	"algorithm/leetcode/ds"
)

func mergeTwoLists(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	// return mergeTwoListsIterate(l1, l2)
	return mergeTwoListsRecursive(l1, l2, &ds.ListNode{})
}

func mergeTwoListsRecursive(l1 *ds.ListNode, l2 *ds.ListNode, res *ds.ListNode) *ds.ListNode {
	if l1 == nil && l2 == nil {
		return res
	} else if l1 == nil {
		res.Next = l2
	} else if l2 == nil {
		res.Next = l1
	} else if l1.Val > l2.Val {
		res.Next = mergeTwoListsRecursive(l1, l2.Next, res)
	} else {
		res.Next = mergeTwoListsRecursive(l1.Next, l2, res)
	}
	return res.Next
}

func mergeTwoListsIterate(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	res := &ds.ListNode{}
	head := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			head.Next = l2
			break
		} else if l2 == nil {
			head.Next = l1
			break
		} else if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
	}
	return res.Next
}
