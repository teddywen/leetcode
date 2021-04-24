/*
https://leetcode-cn.com/problems/ugly-number/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		n = 14
	)
	fmt.Printf("%t", isUgly(n))
}

func isUgly(n int) bool {
	var (
		dp = make([]bool, n+1)
	)
	dp[0] = true
	dp[1] = true
	for i := 2; i <= n; i++ {
		var ugly bool
		for _, v := range []int{2, 3, 5} {
			if i%v == 0 && dp[i/v] {
				ugly = true
			}
		}
		dp[i] = ugly
	}
	return dp[n]
}
