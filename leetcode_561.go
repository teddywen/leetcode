/*
https://leetcode-cn.com/problems/array-partition-i/
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		nums = []int{6, 2, 6, 5, 1, 2}
	)
	fmt.Printf("%d", arrayPairSum(nums))
}

func arrayPairSum(nums []int) int {
	var (
		res int
	)

	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}
