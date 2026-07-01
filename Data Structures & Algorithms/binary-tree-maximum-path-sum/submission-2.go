/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    // idea will be to go down to the furthest leaf node.
	// we want to find the 'maxSum' for left/right.
	// we then return max(left + current, right + current, current, 0
	mx := math.MinInt

	var maxSum func(root *TreeNode) int
	maxSum = func(root *TreeNode) int {
		l := 0
		r := 0
		if root.Left != nil {
			l = maxSum(root.Left)
		}
		if root.Right != nil {
			r = maxSum(root.Right)
		}
		sum := root.Val
		if l > 0 {
			sum += l
		}
		if r > 0 {
			sum += r
		}
		if sum > mx {
			mx = sum
		}
		return max(root.Val + max(0, l, r), 0)
	}

	maxSum(root)
	return mx
}
