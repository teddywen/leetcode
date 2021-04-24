/*
https://leetcode-cn.com/problems/russian-doll-envelopes/
动规解法
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		nums = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	)
	fmt.Printf("%d", maxEnvelopes(nums))
}

// 左界二分
func maxEnvelopes(envelopes [][]int) int {
	var (
		sorted EnvSort = envelopes
		piles [][]int
	)
	sort.Sort(sorted)

	for _, envelope := range sorted {

		if len(piles) == 0 {
			piles = append(piles, envelope)
			continue
		}

		left, right := 0, len(piles)
		for left < right {
			mid := left + (right - left) / 2
			if piles[mid][0] < envelope[0] && piles[mid][1] < envelope[1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == len(piles) {
			piles = append(piles, envelope)
		} else {
			piles[left] = envelope
		}
	}

	return len(piles)
}

// 动态规划
//func maxEnvelopes(envelopes [][]int) int {
//	var (
//		sorted EnvSort = envelopes
//		// 定义 dp[i] = x
//		// x 为 从起始到第 i 个元素的最多套娃数
//		// 方程 dp[i] = dp[j]+1  j 为从i->0途中第一个能够套进当前信封的最多套娃数
//		// base case: dp[0] = 0
//		dp  = make([]int, len(sorted)+1)
//		res int
//	)
//	sort.Sort(sorted)
//
//	for i := 1; i < len(dp); i++ {
//		dp[i] = 1
//		for j := i - 1; j > 0; j-- {
//			if sorted[j-1][0] < sorted[i-1][0] && sorted[j-1][1] < sorted[i-1][1] {
//				dp[i] = max354(dp[i], dp[j]+1)
//			}
//		}
//		res = max354(res, dp[i])
//	}
//
//	return res
//}

func max354(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type EnvSort [][]int

func (es EnvSort) Len() int {
	return len(es)
}

func (es EnvSort) Less(i, j int) bool {
	if es[i][0] < es[j][0] {
		return true
	}
	if es[i][0] == es[j][0] && es[i][1] > es[j][1] {
		return true
	}
	return false
}

func (es EnvSort) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}
