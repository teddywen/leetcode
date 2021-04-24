/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 */
package main

import "fmt"

func main() {
	var (
		nums = []int{7,1,5,3,6,4}
	)
	fmt.Printf("%d", maxProfit(nums))
}

// 差值解法（自创）
//func maxProfit(prices []int) int {
//	delta := make([]int, len(prices))
//	for i:=1; i<len(prices); i++ {
//		delta[i] = prices[i] - prices[i-1]
//	}
//	var res, sum int
//	for _, d := range delta {
//		if sum + d > 0 {
//			sum += d
//			if sum > res {
//				res = sum
//			}
//		} else {
//			sum = 0
//		}
//	}
//	return res
//}

// 历史最低点解法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		min = prices[0]
		res int
	)

	for i:=1; i<len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			delta := prices[i] - min
			if delta > res {
				res = delta
			}
		}
	}
	return res
}