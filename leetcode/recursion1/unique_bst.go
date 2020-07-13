package recursion1

import (
	"algorithm/leetcode/structure"
)

func generateTrees(n int) []*structure.TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreeRecursive(1, n)
}

func generateTreeRecursive(start int, end int) []*structure.TreeNode {
	var list []*structure.TreeNode

	if start > end {
		list = append(list, nil)
		return list
	}

	var left, right []*structure.TreeNode
	for i := start; i <= end; i++ {
		left = generateTreeRecursive(start, i-1)
		right = generateTreeRecursive(i+1, end)

		for _, lnode := range left {
			for _, rnode := range right {
				root := &structure.TreeNode{
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
