/*
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
*/
package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		// step. pop and push
		depth++
		p := q
		q = make([]*TreeNode, 0)
		for _, n := range p {
			if n.Left == nil && n.Right == nil {
				return depth
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
	return depth
}
