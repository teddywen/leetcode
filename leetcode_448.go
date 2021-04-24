/*
https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
 */
package main

import "fmt"

func main() {
	var (
		nums = []int{4,3,2,7,8,2,3,1}
	)
	fmt.Printf("%#v", findDisappearedNumbers(nums))
}

// 需要有个hash表记录每个数字对应的统计，最后把hash表里统计为0的索引拿出来就是答案
// 由于题目要求不使用额外空间，因此我们可以复用nums做hash表。
// 为了不影响原值的读写，我们在统计哈希值时在原值上面加n，这样通过v%n依然能拿到原值
func findDisappearedNumbers(nums []int) []int {
	var (
		n = len(nums)
		result []int
	)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		// v < n，说明索引为i的数没被统计过，因此在数组里不存在
		if v <= n {
			result = append(result, i + 1)
		}
	}
	return result
}