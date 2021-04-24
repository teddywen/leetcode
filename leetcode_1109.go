/*
https://leetcode-cn.com/problems/corporate-flight-bookings/
*/
package main

import "fmt"

func main() {
	var (
		bookings = [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
		n        = 5
	)
	fmt.Printf("%#v", corpFlightBookings(bookings, n))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	var (
		diff = make([]int, n)
		res  = make([]int, n)
	)
	for _, booking := range bookings {
		incrementDiff(diff, booking[0]-1, booking[1]-1, booking[2])
	}
	for i := range diff {
		if i == 0 {
			res[i] = diff[i]
		} else {
			res[i] = res[i-1] + diff[i]
		}
	}
	return res
}

func incrementDiff(diff []int, i, j, incr int) {
	diff[i] += incr
	if j+1 < len(diff) {
		diff[j+1] -= incr
	}
}
