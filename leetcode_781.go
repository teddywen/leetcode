/*
https://leetcode-cn.com/problems/rabbits-in-forest/
*/
package main

import "fmt"

func main() {
	var (
		answers = []int{1, 1, 2}
	)
	fmt.Printf("%d", numRabbits(answers))
}

func numRabbits(answers []int) int {
	var (
		counter = make(map[int]int)
		result  int
	)

	for _, same := range answers {
		counter[same+1]++
	}

	for same, count := range counter {
		result += (count + same - 1) / same * same
	}

	return result
}
