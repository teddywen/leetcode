/*
https://leetcode-cn.com/problems/longest-harmonious-subsequence/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{1, 3, 2, 2, 5, 2, 3, 7}
	)
	fmt.Printf("%d", findLHS(nums))
}

func findLHS(nums []int) int {
	var (
		hash = make(map[int]int, len(nums))
		result int
	)
	for _, num := range nums {
		hash[num]++
		if hash[num-1] > 0 && result < hash[num] + hash[num-1] {
			result = hash[num] + hash[num-1]
		}
		if hash[num+1] > 0 && result < hash[num] + hash[num+1] {
			result = hash[num] + hash[num+1]
		}
	}
	return result
}
