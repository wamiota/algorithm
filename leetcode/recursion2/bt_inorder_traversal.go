package recursion2

import "algorithm/leetcode/ds"

func inorderTraversal(root *ds.TreeNode) []int {
	var res []int
	res = inorderTraversalRecursive(root, res)
	return res
}

func inorderTraversalRecursive(node *ds.TreeNode, list []int) []int {
	if node == nil {
		return list
	}

	list = inorderTraversalRecursive(node.Left, list)
	list = append(list, node.Val)
	list = inorderTraversalRecursive(node.Right, list)

	return list
}
