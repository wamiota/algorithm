package recursion2

import "algorithm/leetcode/ds"

func isValidBST(root *ds.TreeNode) bool {
	return isValidBSTRecursive(root, nil, nil)
}

func isValidBSTRecursive(node *ds.TreeNode, lo *int, hi *int) bool {
	if node == nil {
		return true
	}

	if lo != nil && node.Val <= *lo {
		return false
	}
	if hi != nil && node.Val >= *hi {
		return false
	}

	return isValidBSTRecursive(node.Left, lo, &node.Val) &&
		isValidBSTRecursive(node.Right, &node.Val, hi)
}
