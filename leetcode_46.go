/*
https://leetcode-cn.com/problems/permutations/
 */
package main

import "fmt"

func main() {
	nums := []int{5,4,6,2}
	fmt.Printf("%#v", permute(nums))
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	chosen := make(map[int]struct{})
	backtrace(&result, nums, nil, chosen)
	return result
}

func backtrace(result *[][]int, choice []int, path []int, chosen map[int]struct{}) {
	if len(choice) == len(chosen) {
		*result = append(*result, path)
		return
	}

	for _, c := range choice {
		if _, ok := chosen[c]; ok {
			continue
		}
		// 做选择
		chosen[c] = struct{}{}

		// 回溯递归
		backtrace(result, choice, append(path, c), chosen)

		// 撤销选择
		delete(chosen, c)
	}
}
