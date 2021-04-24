/*
https://leetcode-cn.com/problems/grumpy-bookstore-owner/
*/
package main

import "fmt"

func main() {
	var (
		customers = []int{1, 0, 1, 2, 1, 1, 7, 5}
		grumpy    = []int{0, 1, 0, 1, 0, 1, 0, 1}
		X         = 3
	)
	fmt.Printf("%d", maxSatisfied(customers, grumpy, X))
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	var (
		satisfiedSum, maxGoToSatisfied int
	)

	for i := range customers {
		// 统计所有满意的人
		satisfiedSum += customers[i] * (grumpy[i] ^ 1)
		// 用customers存最大前戳和
		customers[i] *= grumpy[i]
		if i > 0 {
			customers[i] += customers[i-1]
		}
		var preSum int
		if i-X >= 0 {
			preSum = customers[i-X]
		}
		if customers[i]-preSum > maxGoToSatisfied {
			maxGoToSatisfied = customers[i] - preSum
		}
	}

	return satisfiedSum + maxGoToSatisfied
}
