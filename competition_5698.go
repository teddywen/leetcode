/*
https://leetcode-cn.com/contest/weekly-contest-231/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{1,-1,1}
		limit = 3
		goal = -4
	)
	fmt.Printf("%d", minElements(nums, limit, goal))
}

func minElements(nums []int, limit int, goal int) int {
	var (
		sum, distance int
	)

	for _, v := range nums {
		sum += v
	}

	distance = int(math.Abs(float64(int(math.Abs(float64(sum - goal))))))
	return int(math.Ceil(float64(distance) / float64(limit)))
}