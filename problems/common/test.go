package main

import (
	"fmt"
)

/*
test case:
"adc"
"dcda"
true

"hello"
"ooolleoooleh"
false

"abc"
"bbbca"
true

"abcdxabcde"
"abcdeabcdx"
true

https://leetcode.com/submissions/detail/267059201/testcase/
false

https://leetcode.com/submissions/detail/267059572/testcase/
true
*/

func main() {
	res := subarraySum([]int{0, 0, 0}, 0)
	// 998001
	fmt.Println(res)
}

// func subarraySum(nums []int, k int) int {
// 	res := 0
// 	// key is sum, value is count
// 	sumMap := make(map[int]bool, 0)
// 	sumMap[0] = true
// 	sum := 0
// 	for i := 0; i < len(nums); i++ {
// 		num := nums[i]
// 		sum += num
// 		if _, ok := sumMap[sum-k]; ok {
// 			res++
// 		}
// 		sumMap[sum] = true
// 	}
// 	return res
// }

func subarraySum(nums []int, k int) int {
	res := 0
	// key is sum, value is count
	sumMap := make(map[int]int, 0)
	sumMap[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		sum += num
		if _, ok := sumMap[sum-k]; ok {
			res += sumMap[sum-k]
		}
		sumMap[sum] += 1
	}
	return res
}
