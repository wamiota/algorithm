package recursion1

import (
	"algorithm/leetcode/structure"
)

func searchBST(root *structure.TreeNode, val int) *structure.TreeNode {
	if root == nil {
		return root
	}

	return searchBSTRecursive(root, val)
}

func searchBSTRecursive(node *structure.TreeNode, val int) *structure.TreeNode {
	if node.Val == val {
		return node
	}

	left := searchBSTRecursive(node.Left, val)
	if left.Val == val {
		return left
	}

	return searchBSTRecursive(node.Right, val)
}
