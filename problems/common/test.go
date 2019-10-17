package main

import (
	"fmt"
)

/*
test case:
[[1,1,4],[9,4,9],[9,1,9],[2,3,5],[4,1,5],[10,4,5]]
33
false

*/

func main() {
	res := findDuplicates([]int{5, 4, 6, 7, 9, 3, 10, 9, 5, 6})
	// 998001
	fmt.Println(res)
}
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	i := 0
	for i < len(nums) {
		num := nums[i]
		if nums[i] == i+1 {
			i++
			continue
		}
		if num == -1 {
			i++
			continue
		}
		if nums[num-1] == num {
			nums[i] = -1
			res = append(res, num)
			i++
			continue
		}
		nums[num-1], nums[i] = nums[i], nums[num-1]
	}
	return res
}
