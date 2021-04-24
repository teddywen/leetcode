/*
https://leetcode-cn.com/problems/russian-doll-envelopes/
二分搜索解法
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		nums = [][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}
	)
	fmt.Printf("%d", maxEnvelopes(nums))
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	// 对信封从小到大排序
	s := EnvSort{Data: envelopes}
	sort.Sort(s)
	envelopes = s.Data

	var (
		piles = make([][]int, 0) // 堆
		left  = 0
		right = 0
	)

	for _, envelope := range envelopes {
		// 二分法找左界
		left = 0
		right = len(piles)
		for left < right {
			mid := left + (right-left)/2
			if Lte(envelope, piles[mid]) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == len(piles) {
			piles = append(piles, envelope)
		} else {
			piles[left] = envelope
		}
	}

	return len(piles) // 堆数就是最大子序列长度
}

func Lte(x []int, y []int) bool {
	return !(x[0] > y[0] && x[1] > y[1])
}

type EnvSort struct {
	Data [][]int
}

func (e EnvSort) Len() int {
	return len(e.Data)
}
func (e EnvSort) Less(i, j int) bool {
	return e.Data[i][0] < e.Data[j][0] || (e.Data[i][0] == e.Data[j][0] && e.Data[i][1] > e.Data[j][1])
}
func (e EnvSort) Swap(i, j int) {
	e.Data[i], e.Data[j] = e.Data[j], e.Data[i]
}
