/*
https://leetcode-cn.com/problems/number-of-1-bits/
 */
package main

import "fmt"

func main() {
	var (
		nums = []int{2,7,11,15}
		target = 9
	)
	fmt.Printf("%#v", twoSum(nums, target))
}

// 位运算优化
// 去掉位中的最后一位，循环的次数即1的数量
func hammingWeight(num uint32) int {
	var (
		result int
	)

	if num == 0 {
		return 0
	}

	for num != 0 {
		num &= num - 1
		result++
	}
	return result
}

// 循环检查
//func hammingWeight(num uint32) int {
//	var (
//		result int
//	)
//
//	for num != 0 {
//		result += int(num & 1)
//		num = num >> 1
//	}
//	return result
//}