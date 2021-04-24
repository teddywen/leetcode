/*
https://leetcode-cn.com/problems/reshape-the-matrix/
*/
package main

import "fmt"

func main() {
	var (
		nums = [][]int{{1, 2}, {3, 4}}
		r    = 1
		c    = 4
	)
	fmt.Printf("%+v", matrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	var (
		rr   = len(nums)
		cc   = len(nums[0])
		i, j int
	)
	if rr*cc != r*c {
		return nums
	}

	var res = make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	for ii := 0; ii < rr; ii++ {
		for jj := 0; jj < cc; jj++ {
			res[i][j] = nums[ii][jj]
			j++
			if j == c {
				i++
				j = 0
			}
		}
	}

	return res
}
