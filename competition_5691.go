/*
https://leetcode-cn.com/contest/weekly-contest-230/problems/equal-sum-arrays-with-minimum-number-of-operations/
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		nums1 = []int{5, 2, 1, 5, 2, 2, 2, 2, 4, 3, 3, 5}
		nums2 = []int{1, 4, 5, 5, 6, 3, 1, 3, 3}
	)
	fmt.Printf("%#v", minOperations5691(nums1, nums2))
}

func minOperations5691(nums1 []int, nums2 []int) int {
	var (
		sum1, sum2, delta            int
		minus1, plus1, minus2, plus2 []int
	)
	for _, n := range nums1 {
		sum1 += n
		minus1 = append(minus1, n-1)
		plus1 = append(plus1, 6-n)
	}
	for _, n := range nums2 {
		sum2 += n
		minus2 = append(minus2, n-1)
		plus2 = append(plus2, 6-n)
	}

	if sum1 == sum2 {
		return 0
	}

	delta = sum1 - sum2

	minus1 = append(minus1, plus2...)
	minus2 = append(minus2, plus1...)

	sort.Sort(sort.Reverse(sort.IntSlice(minus1)))
	sort.Sort(sort.Reverse(sort.IntSlice(minus2)))

	var minus []int
	var opNum int
	if delta > 0 {
		minus = minus1
	} else {
		minus = minus2
	}
	delta = int(math.Abs(float64(delta)))
	for _, d := range minus {
		opNum++
		delta -= d
		if delta <= 0 {
			return opNum
		}
	}

	return -1
}
