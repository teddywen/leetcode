/*
https://leetcode-cn.com/contest/weekly-contest-230/problems/count-items-matching-a-rule/
 */
package main

import "fmt"

func main() {
	var (
		items = [][]string{{"phone","blue","pixel"},{"computer","silver","phone"},{"phone","gold","iphone"}}
		ruleKey = "type"
		ruleValue = "phone"
	)
	fmt.Printf("%#v", countMatches(items, ruleKey, ruleValue))
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var (
		checkIdx, cnt int
	)
	switch ruleKey {
	case "type":
		checkIdx = 0
	case "color":
		checkIdx = 1
	case "name":
		checkIdx = 2
	}

	for _, item := range items {
		if item[checkIdx] == ruleValue {
			cnt++
		}
	}

	return cnt
}