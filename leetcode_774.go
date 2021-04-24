/*
https://leetcode-cn.com/problems/subarray-sum-equals-k/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{1, 1, 1}
		k    = 2
	)
	fmt.Printf("%d", subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	var (
		preSum = map[int]int{0: 1}
		curSum, res int
	)
	for _, v := range nums {
		curSum += v
		lastSum := curSum - k
		if cnt, ok := preSum[lastSum]; ok {
			res += cnt
		}
		preSum[curSum] += 1
	}
	return res
}
