/*
https://leetcode-cn.com/problems/increasing-subsequences/
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		nums = []int{1, 3, 5, 7}
	)
	fmt.Printf("%v", findSubsequences(nums))
}

// 暴力求解
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)

	memo := make(map[string]struct{})
	var inMemo = func(seq []int) bool {
		key := ""
		for _, s := range seq {
			key += " " + strconv.Itoa(s)
		}
		_, ok := memo[key]
		return ok
	}
	var setMemo = func(seq []int) {
		key := ""
		for _, s := range seq {
			key += " " + strconv.Itoa(s)
		}
		memo[key] = struct{}{}
	}
	var putRes = func(seq []int) {
		copyOf := make([]int, len(seq))
		copy(copyOf, seq)
		res = append(res, seq)
	}

	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		seq := []int{nums[i]}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] >= tmp {
				tmp = nums[j]
				seq = append(seq, tmp)
				if !inMemo(seq) {
					putRes(seq)
					setMemo(seq)
				}
			}
		}
		for k := 0; k < len(seq); k++ {
			for l := k + 1; l < len(seq); l++ {

			}
		}
	}
	return res
}
