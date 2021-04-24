/*
https://leetcode-cn.com/contest/weekly-contest-230/problems/closest-dessert-cost/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		baseCosts    = []int{1, 7}
		toppingCosts = []int{3, 4}
		target       = 10
	)
	fmt.Printf("%d", closestCost(baseCosts, toppingCosts, target))
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	return 0
}
