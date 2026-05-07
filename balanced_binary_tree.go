package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isBalanced(root *TreeNode) bool {
	_, balanced := checkBalance(root)
	return balanced
}

func checkBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftBalanced := checkBalance(root.Left)
	rightDepth, rightBalanced := checkBalance(root.Right)

	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	if abs(leftDepth-rightDepth) > 1 {
		return 0, false
	}

	return 1 + max(leftDepth, rightDepth), true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}