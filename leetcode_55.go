/*
https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/
*/
package main

import "math"

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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftDepth, rightDepth int
	if root.Left != nil {
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDepth = maxDepth(root.Right)
	}
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
