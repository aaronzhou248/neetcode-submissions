/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 import "slices"

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var pathP []*TreeNode
	var pathQ []*TreeNode
    var bfs func (current *TreeNode, path []*TreeNode)
	bfs = func (current *TreeNode, path []*TreeNode) {
		path = append(path, current)

		if current == p {
			pathP = slices.Clone(path)
		}
		if current == q {
			pathQ = slices.Clone(path)
		}

		if current.Left != nil {
			bfs(current.Left, path)
		}
		if current.Right != nil {
			bfs(current.Right, path)
		}
	}

	bfs(root, []*TreeNode{})

	lenP, lenQ := len(pathP), len(pathQ)
	output := root
	for i := range(min(lenP, lenQ) - 1) {
		if pathP[i + 1] == pathQ[i + 1] {
			output = pathP[i + 1]
		}
	}
	return output
}
