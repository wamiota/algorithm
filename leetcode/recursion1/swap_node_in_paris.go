package recursion1

import (
	"algorithm/leetcode/structure"
)

func swapPairs(head *structure.ListNode) *structure.ListNode {
	return recursiveSwapPairs(head)
}

func recursiveSwapPairs(curNode *structure.ListNode) *structure.ListNode {
	if curNode == nil || curNode.Next == nil {
		return curNode
	}

	temp := curNode.Val
	curNode.Val = curNode.Next.Val
	curNode.Next.Val = temp

	recursiveSwapPairs(curNode.Next.Next)

	return curNode
}
