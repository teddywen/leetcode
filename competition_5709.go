/*
https://leetcode-cn.com/contest/weekly-contest-233/problems/maximum-ascending-subarray-sum/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		nums = []int{12, 17, 15, 13, 10, 11, 12}
	)
	fmt.Printf("%d", maxAscendingSum(nums))
}

func maxAscendingSum(nums []int) int {
	var (
		acc, last, result int
	)
	for _, n := range nums {
		if last < n {
			acc += n
		} else {
			acc = n
		}

		if result < acc {
			result = acc
		}

		last = n
	}

	return result
}
