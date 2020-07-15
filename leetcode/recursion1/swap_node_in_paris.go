package recursion1

import (
	"algorithm/leetcode/ds"
)

func swapPairs(head *ds.ListNode) *ds.ListNode {
	return recursiveSwapPairs(head)
}

func recursiveSwapPairs(curNode *ds.ListNode) *ds.ListNode {
	if curNode == nil || curNode.Next == nil {
		return curNode
	}

	temp := curNode.Val
	curNode.Val = curNode.Next.Val
	curNode.Next.Val = temp

	recursiveSwapPairs(curNode.Next.Next)

	return curNode
}
