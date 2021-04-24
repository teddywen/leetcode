/*
https://leetcode-cn.com/problems/subarrays-with-k-different-integers/
*/
package main

import "fmt"

func main() {
	var (
		A = []int{1, 2, 1, 2, 3}
		K = 2
	)
	fmt.Printf("%d", subarraysWithKDistinct(A, K))
}

// 滑动窗口解法
func subarraysWithKDistinct(A []int, K int) int {
	var (
		n                             = len(A)
		l1, l2, r, tot1, tot2, result int
		num1                          = make([]int, n+1)
		num2                          = make([]int, n+1)
	)

	for r < len(A) {
		// 滑动窗口往右伸展
		c := A[r]
		r++

		// 处理数据
		if num1[c] == 0 {
			tot1++
		}
		num1[c]++
		if num2[c] == 0 {
			tot2++
		}
		num2[c]++

		// l1窗口收缩
		for tot1 > K && l1 < n {
			lc := A[l1]
			l1++
			if num1[lc] > 0 {
				num1[lc]--
				if num1[lc] == 0 {
					tot1--
				}
			}
		}
		// l2窗口收缩
		for tot2 > K-1 && l2 < n {
			lc := A[l2]
			l2++
			if num2[lc] > 0 {
				num2[lc]--
				if num2[lc] == 0 {
					tot2--
				}
			}
		}

		if l2-l1 > 0 {
			result += l2 - l1
		}
	}

	return result
}
