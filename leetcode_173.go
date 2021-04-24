/*
https://leetcode-cn.com/problems/binary-search-tree-iterator/
*/
package main

import "fmt"

func main() {
	node1 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 3}
	node15 := &TreeNode{Val: 15}
	node9 := &TreeNode{Val: 9}
	node20 := &TreeNode{Val: 20}
	node1.Left = node3
	node1.Right = node15
	node15.Left = node9
	node15.Right = node20
	it := Constructor(node1)
	fmt.Printf("%d,%d,%t,%d,%t,%d,%t,%d,%t",
		it.Next(), it.Next(), it.HasNext(), it.Next(), it.HasNext(),it.Next(), it.HasNext(), it.Next(), it.HasNext())
}

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
type BSTIterator struct {
	NodeStack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{NodeStack: []*TreeNode{root}}
}

func (this *BSTIterator) Next() int {
	popNode := this.NodeStack[len(this.NodeStack)-1]
	for popNode.Left != nil {
		this.NodeStack = append(this.NodeStack, popNode.Left)
		popNode, popNode.Left = popNode.Left, nil
	}

	popNode = this.NodeStack[len(this.NodeStack)-1]
	this.NodeStack = this.NodeStack[:len(this.NodeStack)-1]
	//result := popNode.Val

	if popNode.Right != nil {
		this.NodeStack = append(this.NodeStack, popNode.Right)
	}

	return popNode.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.NodeStack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
