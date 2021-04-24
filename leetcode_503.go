/*
https://leetcode-cn.com/problems/daily-temperatures/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{1, 2, 1}
	)
	fmt.Printf("%#v", nextGreaterElements(nums))
}

func nextGreaterElements(nums []int) []int {
	var (
		monotonousStack = make([]int, 0)
	)

	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(monotonousStack) > 0 && nums[i%len(nums)] >= monotonousStack[len(monotonousStack)-1] {
			monotonousStack = monotonousStack[:len(monotonousStack)-1]
		}
		v := nums[i%len(nums)]
		if i < len(nums) {
			if len(monotonousStack) == 0 {
				nums[i] = -1
			} else {
				nums[i] = monotonousStack[len(monotonousStack)-1]
			}
		}
		monotonousStack = append(monotonousStack, v)
	}

	return nums
}

//func nextGreaterElements(nums []int) []int {
//	var (
//		monotonousStack = make([]int, 0)
//	)
//
//	for i := len(nums)*2 - 1; i >= 0; i-- {
//		v := nums[i%len(nums)]
//		for len(monotonousStack) > 0 && v >= monotonousStack[len(monotonousStack)-1] {
//			monotonousStack = monotonousStack[:len(monotonousStack)-1]
//		}
//		if i < len(nums) {
//			if len(monotonousStack) == 0 {
//				nums[i] = -1
//			} else {
//				nums[i] = monotonousStack[len(monotonousStack)-1]
//			}
//		}
//		monotonousStack = append(monotonousStack, v)
//	}
//
//	return nums
//}
