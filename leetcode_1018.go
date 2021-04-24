/*
https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{0, 1, 1}
	)
	fmt.Printf("%#v", prefixesDivBy5(nums))
}

func prefixesDivBy5(A []int) []bool {
	var prefix int
	var res []bool
	for _, a := range A {
		prefix = ((prefix << 1) + a) % 5
		res = append(res, prefix == 0)
	}
	return res
}
