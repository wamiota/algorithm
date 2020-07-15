package recursion1

import (
	"algorithm/leetcode/ds"
)

func generateTrees(n int) []*ds.TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreeRecursive(1, n)
}

func generateTreeRecursive(start int, end int) []*ds.TreeNode {
	var list []*ds.TreeNode

	if start > end {
		list = append(list, nil)
		return list
	}

	var left, right []*ds.TreeNode
	for i := start; i <= end; i++ {
		left = generateTreeRecursive(start, i-1)
		right = generateTreeRecursive(i+1, end)

		for _, lnode := range left {
			for _, rnode := range right {
				root := &ds.TreeNode{
					Val: i,
				}
				root.Left = lnode
				root.Right = rnode
				list = append(list, root)
			}
		}
	}
	return list
}
