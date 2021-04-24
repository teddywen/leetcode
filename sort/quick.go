/*
快速排序
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	)
	quickSort(nums, 0, len(nums)-1)
	fmt.Printf("%#v", nums)
}

// 排序[i,j]区间的切片
func quickSort(nums []int, i, j int) {
	if i >= j {
		return
	}

	var (
		left  = i
		right = j
		lToR  = false
	)
	for left < right {
		if lToR {
			// 左往右
			if nums[left] > nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
				lToR = !lToR
				right--
			} else {
				left++
			}
		} else {
			// 右往左
			if nums[right] < nums[left] {
				nums[left], nums[right] = nums[right], nums[left]
				lToR = !lToR
				left++
			} else {
				right--
			}
		}
	}
	quickSort(nums, i, left-1)
	quickSort(nums, left+1, j)
}
