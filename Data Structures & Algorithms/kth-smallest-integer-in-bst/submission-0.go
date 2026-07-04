/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var output *TreeNode
	var bfs func(node *TreeNode, current int) int
	bfs = func(node *TreeNode, current int) int {
		if output != nil {
			return -1
		}
		if node.Left == nil {
			current++
		} else {
			current = bfs(node.Left, current) + 1
		}
		if current == k {
			output = node
		}
		if node.Right != nil {
			current = bfs(node.Right, current)
		}
		return current
	}

	bfs(root, 0)
	return output.Val
}


/*
          6
    3
1       5
  2  4
*/