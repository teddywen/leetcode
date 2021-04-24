/*
归并排序
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	)
	fmt.Printf("%#v", mergeSort(nums))
}

// 排序[i,j]区间的切片
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 分
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	// 并
	var i, j int
	var isLeft bool
	var res []int
	for {
		if isLeft {
			if left[i] <= right[j] {
				res = append(res, left[i])
				i++
				if i == len(left) {
					res = append(res, right[j:]...)
					break
				}
			} else {
				isLeft = !isLeft
			}
		} else {
			if right[j] <= left[i] {
				res = append(res, right[j])
				j++
				if j == len(right) {
					res = append(res, left[i:]...)
					break
				}
			} else {
				isLeft = !isLeft
			}
		}
	}
	return res
}
