func maxPathSum(root *TreeNode) int {
	best := math.MinInt

	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// gain from each side, dropping negatives
		left := max(maxGain(node.Left), 0)
		right := max(maxGain(node.Right), 0)

		// path that peaks HERE, using both children — candidate for the answer
		best = max(best, node.Val+left+right)

		// but only ONE side can extend upward to the parent
		return node.Val + max(left, right)
	}

	maxGain(root)
	return best
}