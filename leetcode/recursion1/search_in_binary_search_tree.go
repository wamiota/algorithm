package recursion1

import (
	"algorithm/leetcode/ds"
)

func searchBST(root *ds.TreeNode, val int) *ds.TreeNode {
	if root == nil {
		return root
	}

	return searchBSTRecursive(root, val)
}

func searchBSTRecursive(node *ds.TreeNode, val int) *ds.TreeNode {
	if node.Val == val {
		return node
	}

	left := searchBSTRecursive(node.Left, val)
	if left.Val == val {
		return left
	}

	return searchBSTRecursive(node.Right, val)
}
