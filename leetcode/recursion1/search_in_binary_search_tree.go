package recursion1

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	return searchBSTRecursive(root, val)
}

func searchBSTRecursive(node *TreeNode, val int) *TreeNode {
	if node.Val == val {
		return node
	}

	left := searchBSTRecursive(node.Left, val)
	if left.Val == val {
		return left
	}

	return searchBSTRecursive(node.Right, val)
}
