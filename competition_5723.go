/*
https://leetcode-cn.com/contest/weekly-contest-235/problems/finding-the-users-active-minutes/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		logs = [][]int{{1,1}, {2,2}, {2,3}}
		k    = 4
	)
	fmt.Printf("%+v", findingUsersActiveMinutes(logs, k))
}

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	var (
		userMinutes = make(map[int]map[int]struct{})
		result = make([]int, k)
	)
	for _, log := range logs {
		id, minute := log[0], log[1]
		if _, ok := userMinutes[id]; !ok {
			userMinutes[id] = make(map[int]struct{})
		}
		userMinutes[id][minute] = struct{}{}
	}
	for _, userMinute := range userMinutes {
		if len(userMinute) > 0 && len(userMinute) <= k {
			result[len(userMinute)-1]++
		}
	}
	return result
}
