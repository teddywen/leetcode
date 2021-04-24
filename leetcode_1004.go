/*
https://leetcode-cn.com/problems/max-consecutive-ones-iii/
 */
package main

import "fmt"

func main() {
	var (
		A = []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
		K = 3
	)
	fmt.Printf("%d", longestOnes(A, K))
}

func longestOnes(A []int, K int) int {
	var (
		left, right, cnt, res int
	)
	for right < len(A) {
		c := A[right]
		right++

		if c == 0 {
			cnt++
		}

		for cnt > K && left < right {
			c := A[left]
			left++

			if c == 0 {
				cnt--
			}
		}

		max := right - left
		if res < max {
			res = max
		}
	}
	return res
}