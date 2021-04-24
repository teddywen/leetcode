/*
插入排序
 */
package main

import "fmt"

func main() {
	var (
		nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	)
	insertion(nums)
	fmt.Printf("%#v", nums)
}

func insertion(nums []int) {
	for i := range nums {
		j := i
		for j > 0 {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				j--
			} else {
				break
			}
		}
	}
}
