/*
https://leetcode-cn.com/problems/move-zeroes/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{0, 1, 0, 3, 12}
	)
	moveZeroes(nums)
	fmt.Printf("%#v", nums)
}

func moveZeroes(nums []int) {
	var (
		i, j int
	)
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
