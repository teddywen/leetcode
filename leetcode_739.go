/*
https://leetcode-cn.com/problems/daily-temperatures/
*/
package main

import "fmt"

func main() {
	var (
		temperatures = []int{73, 74, 75, 71, 69, 72, 76, 73}
	)
	fmt.Printf("%#v", dailyTemperatures(temperatures))
}

func dailyTemperatures(T []int) []int {
	var (
		monotonousStack = make([]int, 0)
		result          = make([]int, len(T))
	)

	for i := len(T) - 1; i >= 0; i-- {
		v := T[i]
		for len(monotonousStack) > 0 && v >= T[monotonousStack[len(monotonousStack)-1]] {
			monotonousStack = monotonousStack[:len(monotonousStack)-1]
		}
		if len(monotonousStack) == 0 {
			result[i] = 0
		} else {
			result[i] = monotonousStack[len(monotonousStack)-1] - i
		}
		monotonousStack = append(monotonousStack, i)
	}

	return result
}
