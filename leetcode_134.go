/*
https://leetcode-cn.com/problems/gas-station/
*/
package main

import "fmt"

func main() {
	var (
		gas  = []int{1, 2, 3, 4, 5}
		cost = []int{3, 4, 5, 1, 2}
	)
	fmt.Printf("%d", canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	var (
		n                = len(gas)
		tail             = 2*n - 1
		left, right, sum int
	)
	for i := range gas {
		cost[i] = gas[i] - cost[i]
	}

	for right < tail {
		c := cost[right%n]
		right++
		sum += c

		for sum < 0 {
			c := cost[left%n]
			left++
			sum -= c
		}

		if right-left == n {
			return left
		}
	}

	return -1
}
