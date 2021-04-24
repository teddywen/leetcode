/*
https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/
*/
package main

import "fmt"

func main() {
	var (
		A = []int{0, 0, 0, 1, 0, 1, 1, 0}
		K = 3
	)
	fmt.Printf("%d", minKBitFlips(A, K))
}

func minKBitFlips(A []int, K int) int {
	if len(A) < K {
		return -1
	}

	var (
		n        = len(A)
		diff     = make([]int, n+1)
		res, rev int
	)
	for i, v := range A {
		rev += diff[i]
		if (v+rev)%2 == 0 {
			if i+K > n {
				return -1
			}
			rev++
			diff[i+K]--
			res++
		}
	}
	return res
}
