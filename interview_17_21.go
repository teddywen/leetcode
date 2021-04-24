/*
https://leetcode-cn.com/problems/volume-of-histogram-lcci/
*/
package main

import "fmt"

func main() {
	var (
		height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	)
	fmt.Printf("%d", trap(height))
}

// 双指针
func trap(height []int) int {
	var (
		i           int
		j           = len(height) - 1
		left, right int
		result      int
	)
	for i <= j {
		if left <= right {
			if left > height[i] {
				result += left - height[i]
			}
			if height[i] > left {
				left = height[i]
			}
			i++
		} else {
			if right > height[j] {
				result += right - height[j]
			}
			if height[j] > right {
				right = height[j]
			}
			j--
		}
	}

	return result
}

// 多次遍历
//func trap(height []int) int {
//	var (
//		leftHeight  = make([]int, len(height))
//		rightHeight = make([]int, len(height))
//		top, result int
//	)
//
//	for i := 0; i < len(height); i++ {
//		leftHeight[i] = top
//		if height[i] > top {
//			top = height[i]
//		}
//	}
//	top = 0
//	for j := len(height) - 1; j >= 0; j-- {
//		rightHeight[j] = top
//		if height[j] > top {
//			top = height[j]
//		}
//	}
//
//	for i := 0; i < len(height); i++ {
//		var min int
//		if leftHeight[i] < rightHeight[i] {
//			min = leftHeight[i]
//		} else {
//			min = rightHeight[i]
//		}
//		if min > height[i] {
//			result += min - height[i]
//		}
//	}
//
//	return result
//}
