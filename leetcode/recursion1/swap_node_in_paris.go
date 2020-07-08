package recursion1

func swapPairs(head *ListNode) *ListNode {
	return recursiveSwapPairs(head)
}

func recursiveSwapPairs(curNode *ListNode) *ListNode {
	if curNode == nil || curNode.Next == nil {
		return curNode
	}

	temp := curNode.Val
	curNode.Val = curNode.Next.Val
	curNode.Next.Val = temp

	recursiveSwapPairs(curNode.Next.Next)

	return curNode
}
