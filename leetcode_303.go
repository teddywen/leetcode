/*
https://leetcode-cn.com/problems/two-sum/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{-2, 0, 3, -5, 2, -1}
	)
	obj := Constructor(nums)
	fmt.Printf("%d\n", obj.SumRange(0, 2))
	fmt.Printf("%d\n", obj.SumRange(2, 5))
	fmt.Printf("%d\n", obj.SumRange(0, 5))
}

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	return NumArray{prefix: prefix}
}

func (na *NumArray) SumRange(i int, j int) int {
	return na.prefix[j+1] - na.prefix[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
