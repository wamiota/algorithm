package recursion1

import "algorithm/leetcode/ds"

func maxDepth(root *ds.TreeNode) int {
	return maxDepthRecursive(root, 0)
}

func maxDepthRecursive(node *ds.TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	left := maxDepthRecursive(node.Left, depth+1)
	right := maxDepthRecursive(node.Right, depth+1)

	if left > right {
		return left
	}
	return right
}
