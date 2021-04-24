/*
https://leetcode-cn.com/problems/two-sum/
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

func twoSum(nums []int, target int) []int {
	mem := make(map[int]int)
	for i, num := range nums {
		remain := target - num
		if remainIdx, ok := mem[remain]; ok {
			return []int{remainIdx, i}
		}
		mem[num] = i
	}
	return []int{}
}
