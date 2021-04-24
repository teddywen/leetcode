/*
https://leetcode-cn.com/problems/fibonacci-number/
*/
package main

import "fmt"

func main() {
	var (
		n = 4
	)
	fmt.Printf("%d", fib(n))
}

// 递归
//func fib(n int) int {
//	if n < 2 {
//		return n
//	}
//	return fib(n - 1) + fib(n - 2)
//}

// 备忘录
//var memo = make(map[int]int)
//func fib(n int) int {
//	if n < 2 {
//		return n
//	}
//
//	if r, ok := memo[n]; ok {
//		return r
//	}
//
//	res := fib(n - 1) + fib(n - 2)
//	memo[n] = res
//
//	return res
//}

// 递推
//func fib(n int) int {
//	if n < 2 {
//		return n
//	}
//
//	dp := make([]int, n)
//	for i := 3; i <= n; i++ {
//		dp[i] = dp[i-1] + dp[i-2]
//	}
//	return dp[1]
//}

// 递推+空间压缩
func fib(n int) int {
	if n < 2 {
		return n
	}

	dp := []int{1, 1}
	for i := 3; i <= n; i++ {
		last := dp[1] + dp[0]
		dp[0] = dp[1]
		dp[1] = last
	}
	return dp[1]
}
