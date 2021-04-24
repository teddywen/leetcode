/*
https://leetcode-cn.com/problems/flatten-nested-list-iterator/
*/
package main

import "fmt"

func main() {
	var (
		nums   = []int{2, 7, 11, 15}
		target = 9
	)
	fmt.Printf("%#v", twoSum(nums, target))
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {}

type NestedIterator struct {
	list   []int
	cursor int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var (
		list []int
	)
	var dfs func(nested *NestedInteger)
	dfs = func(nested *NestedInteger) {
		if nested.IsInteger() {
			list = append(list, nested.GetInteger())
		} else {
			for _, v := range nested.GetList() {
				dfs(v)
			}
		}
	}
	for _, v := range nestedList {
		dfs(v)
	}
	return &NestedIterator{list: list}
}

func (this *NestedIterator) Next() int {
	this.cursor++
	return this.list[this.cursor-1]
}

func (this *NestedIterator) HasNext() bool {
	return this.cursor < len(this.list)
}
