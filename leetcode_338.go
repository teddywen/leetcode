/*
https://leetcode-cn.com/problems/counting-bits/
*/
package main

import "fmt"

func main() {
	var (
		num = 5
	)
	fmt.Printf("%#v", countBits(num))
}

// 暴力解
// n &= n - 1
func countBits(num int) []int {
	var result = make([]int, num+1)
	for i := 0; i <= num; i++ {
		count, n := 0, i
		for n > 0 {
			n &= n - 1
			count++
		}
		result[i] = count
	}
	return result
}

// 动态规划一、奇偶递推法
// 奇: dp[n] = dp[n-1]+1
// 偶: dp[n] = dp[n/2]
func countBits(num int) []int {
	var (
		// dp定义: dp[n] = x。x为n的二进制数中的 1 的数目
		// dp方程:
		// 		当n为奇数: dp[n] = dp[n-1]+1
		//		解释: 二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的 1。
		//		当n为偶数: dp[n] = dp[n/2]
		//		解释: 二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多。
		// base case: dp[0] = 0
		dp = make([]int, num+1)
	)
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}

// 动态规划二、最高位有效位
// dp[n] = dp[n - highBit] + 1
func countBits(num int) []int {
	var (
		// dp定义: dp[n] = x。x为n的二进制数中的 1 的数目
		// dp方程:
		//		当n&(n-1) == 0时，该数是2的整数次幂，最高为为1，其余为0。记录 highBit=n
		// 		dp[n] = dp[n - highBit] + 1
		//		解释: dp[n]相比去掉最高位的数(n-highBit)的dp要多1, 由于从小到大递推，因此dp[n-highBit]已经有值
		// base case: dp[0] = 0
		dp      = make([]int, num+1)
		highBit int
	)
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			highBit = i
			dp[i] = 1
		} else {
			dp[i] = dp[i-highBit] + 1
		}
	}
	return dp
}

// 动态规划三、最低位有效位
// dp[i] = dp[i>>1] + (i & 1)
func countBits(num int) []int {
	var (
		// dp定义: dp[n] = x。x为n的二进制数中的 1 的数目
		// dp方程:
		// 		当n为奇数: dp[n] = dp[n>>1]+1
		//		解释: 二进制表示中，奇数一定比去掉最低位（右移）的二进制数多1。
		//		当n为偶数: dp[n] = dp[n>>1]
		//		解释: 二进制表示中，偶数一定等价去掉最低位的二进制数。
		// base case: dp[0] = 0
		dp = make([]int, num+1)
	)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}

// 动态规划四、最低设置位
// dp[i] = dp[i&(i-1)] + 1
func countBits(num int) []int {
	var (
		// dp定义: dp[n] = x。x为n的二进制数中的 1 的数目
		// dp方程:
		// 		dp[n] = dp[n&(n-1)]+1
		//		解释: dp[n]的1的树木相当于 n去掉最后一个低位1的二进制数的数目+1
		// base case: dp[0] = 0
		dp = make([]int, num+1)
	)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}
