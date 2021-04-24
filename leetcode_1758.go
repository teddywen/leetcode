/*
https://leetcode-cn.com/problems/two-sum/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		s = "10010100"
	)
	fmt.Printf("%d", minOperations(s))
}

// 贪心
func minOperations(s string) int {
	if len(s) < 2 {
		return 0
	}

	var (
		last  = s[0] - '0'
		last2 = last ^ 1
		res1  = 0
		res2  = 1
		res   int
	)

	for i := 1; i < len(s); i++ {
		c := s[i] - '0'
		if last == c {
			res1++
			c ^= 1
		}
		last = c

		c2 := s[i] - '0'
		if last2 == c2 {
			res2++
			c2 ^= 1
		}
		last2 = c2
	}

	res = res1
	if res2 < res {
		res = res2
	}
	return res
}
