/*
https://leetcode-cn.com/contest/weekly-contest-229/problems/maximum-score-from-performing-multiplication-operations/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		nums        = []int{-5, -3, -3, -2, 7, 1}
		multipliers = []int{-10, -5, 3, 4, 6}
	)
	fmt.Printf("%d", maximumScore(nums, multipliers))
}

func maximumScore(nums []int, multipliers []int) int {
	return dfs5687(0, len(nums)-1, 0, nums, multipliers, make(map[string]int))
}

func dfs5687(i, j, k int, nums, multipliers []int, memo map[string]int) int {
	if k == len(multipliers) {
		return 0
	}
	if res, ok := memo[memoKey(i, j, k)]; ok {
		return res
	}
	res := max5687(
		dfs5687(i, j-1, k+1, nums, multipliers, memo)+nums[j]*multipliers[k],
		dfs5687(i+1, j, k+1, nums, multipliers, memo)+nums[i]*multipliers[k])
	memo[memoKey(i, j, k)] = res
	return res
}

func memoKey(i, j, k int) string {
	return fmt.Sprintf("%d_%d_%d", i, j, k)
}

func max5687(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
